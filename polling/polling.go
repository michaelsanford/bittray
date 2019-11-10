package polling

/**
 * Note: This only works with the Bitbucket API 1.0
 */

import (
	"encoding/json"
	"fmt"
	"github.com/coreos/go-semver/semver"
	"github.com/getlantern/systray"
	"github.com/michaelsanford/bittray/config"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var dashboard *BitbucketDashboard

// Poll retrieves pull request data from Bitbucket at a given interval
func Poll(c chan DashboardSummary) {
	user, url := config.GetConfig()
	//endpoint := url + "/rest/api/1.0/dashboard/pull-requests?state=OPEN&role=REVIEWER&participantStatus=UNAPPROVED"
	endpoint := url + "/rest/api/1.0/dashboard/pull-requests?state=OPEN&limit=100"

	pass, ok, _ := config.AskPass()
	if !ok {
		systray.Quit()
	}

	ticker := time.NewTicker(time.Second * 3)
	//defer ticker.Stop()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.SetBasicAuth(user, pass)

	go func(items chan DashboardSummary) {
		cBackOffDelay := 10

		for t := range ticker.C {
			fmt.Printf( "%+v\n", t)
			resp, _ := client.Do(req)

			if resp != nil {
				switch resp.StatusCode {
				case 200:
					bodyText, _ := ioutil.ReadAll(resp.Body)
					p := readDashboard(bodyText, resp.StatusCode)
					fmt.Printf("SEND: %+v\n", p)
					items <- p
					cBackOffDelay = 10
					break
				case 401:
				case 403:
					items <- DashboardSummary{StatusCode: uint16(resp.StatusCode)}
					ticker.Stop()
					break
				case 429:
					items <- DashboardSummary{StatusCode: uint16(resp.StatusCode)}
					cBackOffDelay = backOff(cBackOffDelay)
					break
				default:
					items <- DashboardSummary{StatusCode: uint16(resp.StatusCode)}
					break
				}
			}
		}
	}(c)
}

func readDashboard(bodyText []byte, statusCode int) DashboardSummary {
	ds := DashboardSummary{}

	if bodyText != nil {
		user, _ := config.GetConfig()

		err := json.Unmarshal(bodyText, &dashboard)
		if err != nil {
			panic(err)
		}

		for _, pr := range dashboard.Values {
			if pr.Author.User.Name == user {
				// Authored
				ds.CommentCount += uint8(pr.Properties.CommentCount)
				ds.OpenTasks += uint8(pr.Properties.OpenTaskCount)

				if !ds.NeedsWork {
					for _, reviewer := range pr.Reviewers {
						if reviewer.Status == "NEEDS_WORK" {
							ds.NeedsWork = true
							break
						}
					}
				}
				ds.Authored++
			} else {
				// Reviewing
				for _, reviewer := range pr.Reviewers {
					if reviewer.User.Name == user && reviewer.Status == "UNAPPROVED" {
						ds.Reviewing++
					}
				}
			}
		}
	}

	ds.StatusCode = uint16(statusCode)

	fmt.Printf("MAKE: %+v\n", ds)

	return ds
}

// CheckForUpdate queries the GitHub repo's latest release tag for an update
func CheckForUpdate() (available bool, latestTagName string, err error) {
	vCurrent := semver.New(config.CurrentVersionTag)

	resp, err := http.Get(config.GhAPI)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	j, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	latestTagName = gjson.Get(string(j), "tag_name").Str
	vLatest := semver.New(latestTagName)

	available = vCurrent.LessThan(*vLatest)
	return
}

func backOff(backOffDelay int) int {
	time.Sleep(time.Second * time.Duration(backOffDelay))
	return rand.Intn(backOffDelay) + backOffDelay
}

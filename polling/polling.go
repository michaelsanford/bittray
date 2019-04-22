package polling

/**
 * Note: This only works with the Bitbucket API 1.0
 */

import (
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

var pInterval = config.GetPollingInterval()

// Poll retrieves pull request data from Bitbucket at a given interval
func Poll() <-chan int8 {
	items := make(chan int8, 1)

	user, url := config.GetConfig()
	endpoint := url + "/rest/api/1.0/dashboard/pull-requests?state=OPEN&role=REVIEWER&participantStatus=UNAPPROVED"

	pass, ok, _ := config.AskPass()
	if !ok {
		systray.Quit()
	}

	ticker := time.NewTicker(pInterval)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.SetBasicAuth(user, pass)

	go func(items chan int8) {
		cBackOffDelay := int(10)
		for ; true; <-ticker.C {
			resp, _ := client.Do(req)

			if resp != nil {
				if resp.StatusCode == 200 {
					bodyText, _ := ioutil.ReadAll(resp.Body)
					items <- int8(gjson.Get(string(bodyText), "size").Uint())
					cBackOffDelay = 10
				} else if resp.StatusCode == 401 || resp.StatusCode == 403 {
					ticker.Stop()
					items <- int8(-2)
				} else if resp.StatusCode == 429 {
					items <- int8(-3)
					cBackOffDelay = backOff(cBackOffDelay)
				}
			} else {
				items <- int8(-1)
			}
		}
	}(items)

	return items
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

	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	latestTagName = gjson.Get(string(json), "tag_name").Str
	vLatest := semver.New(latestTagName)

	available = vCurrent.LessThan(*vLatest)
	return
}

func backOff(backOffDelay int) int {
	time.Sleep(time.Second * time.Duration(backOffDelay))
	return rand.Intn(backOffDelay) + backOffDelay
}

package polling

/**
 * Note: This only works with the Bitbucket API 1.0
 */

import (
	"github.com/getlantern/systray"
	"github.com/michaelsanford/bittray/config"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"time"
)

const pollIntervalSec = 10

// Poll retrieves pull request data from Bitbucket at a given interval
func Poll() <-chan int8 {
	items := make(chan int8)

	user, url := config.GetConfig()
	endpoint := url + "/rest/api/1.0/dashboard/pull-requests?state=OPEN&role=REVIEWER&participantStatus=UNAPPROVED"

	pass, ok, _ := config.AskPass()
	if !ok {
		systray.Quit()
	}

	ticker := time.NewTicker(pollIntervalSec * time.Second)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.SetBasicAuth(user, pass)

	go func(items chan int8) {
		for ; true; <-ticker.C {
			resp, _ := client.Do(req)

			if resp != nil && resp.StatusCode == 200 {
				bodyText, _ := ioutil.ReadAll(resp.Body)
				items <- int8(gjson.Get(string(bodyText), "size").Uint())
			} else {
				items <- int8(-1)
			}
		}
	}(items)

	return items
}

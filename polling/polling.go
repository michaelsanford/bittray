package polling

/**
 * Note: This only works with the Bitbucket API 1.0
 */

import (
	"github.com/michaelsanford/bittray/credentials"
	"github.com/michaelsanford/systray"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"time"
)

const pollIntervalSec = 10

// Poll retrieves pull request data from Bitbucket at a given interval
func Poll() <-chan uint8 {
	items := make(chan uint8)

	user, url := credentials.GetConfig()
	endpoint := url + "/rest/api/1.0/dashboard/pull-requests?state=OPEN&role=REVIEWER&participantStatus=UNAPPROVED"

	pass, ok, _ := credentials.AskPass()
	if !ok {
		systray.Quit()
	}

	ticker := time.NewTicker(pollIntervalSec * time.Second)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.SetBasicAuth(user, pass)

	go func(items chan uint8) {
		for ; true; <-ticker.C {
			resp, _ := client.Do(req)
			bodyText, _ := ioutil.ReadAll(resp.Body)
			items <- uint8(gjson.Get(string(bodyText), "size").Uint())
		}
	}(items)

	return items
}

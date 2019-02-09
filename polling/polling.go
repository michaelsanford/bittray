package polling

/**
 * Note: This only works with the Bitbucket API 1.0
 */

import (
	"fmt"
	"github.com/michaelsanford/bittray/credentials"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"time"
)

type PullRequest struct {
	label string
	link  string
}

func Poll() <-chan []PullRequest {
	items := make(chan []PullRequest)

	user, pass, url := credentials.GetCred()
	endpoint := url + "rest/api/1.0/dashboard/pull-requests?state=OPEN&role=REVIEWER"

	ticker := time.NewTicker(10 * time.Second)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.SetBasicAuth(user, pass)

	go func(items chan []PullRequest) {
		for ; true; <-ticker.C {
			resp, _ := client.Do(req)
			bodyText, _ := ioutil.ReadAll(resp.Body)
			s := extract(string(bodyText))
			items <- s
		}
	}(items)

	return items
}

func extract(json string) []PullRequest {

	var prs []PullRequest

	size := uint8(gjson.Get(json, "size").Uint())

	if size > 0 {

		prs = make([]PullRequest, size-1)

		authors := gjson.Get(json, "values.#.author.user.name").Array()
		//names := gjson.Get(json, "values.#.author.user.displayName").Array()
		links := gjson.Get(json, "values.#.links.self.0.href").Array()
		titles := gjson.Get(json, "values.#.title").Array()
		projects := gjson.Get(json, "values.#.fromRef.repository.project.key").Array()

		for i := uint8(0); i < size; i++ {
			prs = append(prs, PullRequest{
				label: fmt.Sprintf("[%s] %s: %s", projects[i].Str, authors[i].Str, titles[i].Str[0:30]),
				link:  links[i].Str})
		}
	}

	return prs
}

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

func Poll() <-chan map[uint8]PullRequest {
	json := make(chan map[uint8]PullRequest)

	user, pass, u := credentials.GetCred()
	endpoint := u + "rest/api/1.0/dashboard/pull-requests?state=OPEN&role=REVIEWER"

	ticker := time.NewTicker(10 * time.Second)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.SetBasicAuth(user, pass)

	go func() {
		for ; true; <-ticker.C {
			resp, _ := client.Do(req)
			bodyText, _ := ioutil.ReadAll(resp.Body)
			s := extract(string(bodyText))
			fmt.Println(s)
			json <- s
		}
	}()

	return json
}

func extract(json string) map[uint8]PullRequest {

	prs := make(map[uint8]PullRequest)

	size := uint8(gjson.Get(json, "size").Uint())

	if size > 0 {
		authors := gjson.Get(json, "values.#.author.user.name").Array()
		//names := gjson.Get(json, "values.#.author.user.displayName").Array()
		links := gjson.Get(json, "values.#.links.self.0.href").Array()
		titles := gjson.Get(json, "values.#.title").Array()
		projects := gjson.Get(json, "values.#.fromRef.repository.project.key").Array()

		for i := uint8(0); i < size; i++ {
			prs[i] = PullRequest{
				label: fmt.Sprintf("[%s] %s: %s", projects[i].Str, authors[i].Str, titles[i].Str[0:30]),
				link:  links[i].Str}
		}
	}

	return prs
}

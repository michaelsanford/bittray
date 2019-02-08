package polling

/**
Note: This only works with the Bitbucket API 1.0
*/

import (
	"fmt"
	"github.com/michaelsanford/bittray/credentials"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type PullRequest struct {
	author  string
	link    url.URL
	title   string
	updated string
}

func Poll() <-chan string {
	json := make(chan string)

	user, pass, u := credentials.GetCred()
	endpoint := u + "rest/api/1.0/dashboard/pull-requests?state=open"

	ticker := time.NewTicker(5 * time.Second)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.SetBasicAuth(user, pass)

	// TODO Remove Debug log
	fmt.Println(req)

	go func() {
		for t := range ticker.C {
			resp, err := client.Do(req)

			if err != nil {
				log.Fatal(err)
			}

			bodyText, err := ioutil.ReadAll(resp.Body)

			fmt.Println(string(bodyText))
			fmt.Println("Tick at", t)

			json <- string(bodyText)
		}
	}()

	return json
}

// TODO Implement this logic in #8
//func Test() {
//var count int

// TODO More elegant way of extracting this
//authors := gjson.Get(dashboard, "values.#.author.user.name").Array()
//names := gjson.Get(dashboard, "values.#.author.user.displayName").Array()
//links := gjson.Get(dashboard, "values.#.links.self.0.href").Array()
//titles := gjson.Get(dashboard, "values.#.title").Array()
//projects := gjson.Get(dashboard, "values.#.fromRef.repository.project.key").Array()

//if len(authors) == len(names) && len(authors) == len(links) {
//	count = len(authors)
//} else {
//	panic("Inconsistent data received from BitBucket API")
//}

// make(map[]PullRequest)
//for i := 0; i < count; i++ {
//	fmt.Println(
//		fmt.Sprintf("Pull Request \n [by] %s (%s) \n [Title]: %s\n [slug]: %s\n %s",
//			names[i], authors[i], titles[i], projects[i], links[i]))
//}
//}

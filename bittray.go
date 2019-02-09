package main

import (
	"fmt"
	"github.com/michaelsanford/bittray/console"
	"github.com/michaelsanford/bittray/credentials"
	"github.com/michaelsanford/bittray/polling"
	"github.com/michaelsanford/bittray/tray"
)

func main() {
	user, pass, url := credentials.GetCred()

	if user == "" || pass == "" || url == "" {
		credentials.DestroyCred()
		credentials.AskCred()
	}

	console.Hide()

	c := polling.Poll()

	go func() {
		for {
			select {
			case <-c:
				// TODO Why is this a pointer? 
				fmt.Println(c)
			}
		}
	}()

	tray.Run()
}

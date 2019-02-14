package main

import (
	"github.com/michaelsanford/bittray/credentials"
	"github.com/michaelsanford/bittray/tray"
)

func main() {
	var ok bool

	user, pass, url := credentials.GetCred()

	if user == "" || pass == "" || url == "" {
		credentials.DestroyCred()
		ok = credentials.AskCred()
		if !ok {
			return
		}
	}

	tray.Run()
}

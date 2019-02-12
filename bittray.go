package main

import (
	"github.com/michaelsanford/bittray/console"
	"github.com/michaelsanford/bittray/credentials"
	"github.com/michaelsanford/bittray/tray"
)

func main() {
	user, pass, url := credentials.GetCred()

	if user == "" || pass == "" || url == "" {
		credentials.DestroyCred()
		credentials.AskCred()
	}

	console.Hide()

	tray.Run()
}

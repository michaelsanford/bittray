package main

import (
	"github.com/michaelsanford/bittray/console"
	"github.com/michaelsanford/bittray/credentials"
	"github.com/michaelsanford/bittray/tray"
)

func main() {
	auth := credentials.GetCred()
	if auth == (credentials.Auth{}) {
		credentials.AskCred()
	}

	console.Hide()

	tray.Run()
}

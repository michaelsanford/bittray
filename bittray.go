package main

import (
	"github.com/michaelsanford/bittray/credentials"
	"github.com/michaelsanford/bittray/tray"
)

func main() {
	var ok bool

	user, url := credentials.GetConfig()

	if user == "" || url == "" {
		credentials.DestroyConfig()
		ok = credentials.AskConfig()
		if !ok {
			return
		}
	}

	tray.Run()
}

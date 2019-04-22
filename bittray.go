//+build windows

package main

import (
	"github.com/michaelsanford/bittray/config"
	"github.com/michaelsanford/bittray/tray"
)

func main() {
	var ok bool

	user, url := config.GetConfig()

	if user == "" || url == "" {
		config.DestroyConfig()
		ok = config.AskConfig()
		if !ok {
			return
		}
	}

	tray.Run()
}

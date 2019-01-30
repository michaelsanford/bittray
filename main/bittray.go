package main

import (
	"../credentials"
	"github.com/michaelsanford/bittray/tray"
)

func main() {
	// TODO panic: assignment to entry in nil map
	tray.OnReady()

	// Test Windows Credential Manager
	credentials.StoreCred("http://bitbucket.org", "username", "1234")
	credentials.GetCred()
}

package main

import (
	"github.com/michaelsanford/bittray/credentials"
	"github.com/michaelsanford/bittray/tray"
)

func main() {
	// Test Windows Credential Manager
	credentials.StoreCred("http://bitbucket.org", "username", "1234")
	credentials.GetCred()

	// Initialize the system tray
	tray.Run()
}

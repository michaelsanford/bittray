package tray

import (
	"fmt"
	"github.com/getlantern/systray"
	"io/ioutil"
	"os"
	"path/filepath"
)

func OnReady() {
	systray.SetIcon(getIcon("checkmark.ico"))
	//systray.SetTitle("Logging in...")
	//systray.SetTooltip("Polling for Pull Requests")
}

func OnExit() {
	// TODO
}

func getIcon(fileName string) []byte {
	cwd, _ := os.Getwd()
	path := filepath.Join(cwd, "assets", fileName)
	fmt.Println(path)

	iconBytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
	return iconBytes
}

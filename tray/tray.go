package tray

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/michaelsanford/bittray/credentials"
	"github.com/michaelsanford/bittray/tray/assets/checkmark"
)

func Run() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(iconCheckmark.Data)
	systray.SetTitle("Bittray")
	systray.SetTooltip("Launched!")

	auth := credentials.GetCred()

	checkOneTwo := systray.AddMenuItem("Check 1..2..", "Test Menu Item")
	mCredential := systray.AddMenuItem("Start Polling", "Begin looping for new PRs")
	mQuit := systray.AddMenuItem("Quit", "Quit bittray")

	go func() {
		for {
			select {
			case <-checkOneTwo.ClickedCh:
				fmt.Println("Clicked Check 1..2..")
			case <-mCredential.ClickedCh:
				fmt.Println(fmt.Sprintf("%s:%s@%s", auth["user"], auth["pass"], auth["url"]))
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// TODO
	fmt.Println("Exiting!")
}

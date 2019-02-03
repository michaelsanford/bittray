package tray

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/michaelsanford/bittray/tray/assets/checkmark"
)

func Run() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(iconCheckmark.Data)
	systray.SetTitle("BitTray")
	systray.SetTooltip("BitTray!")

	checkOneTwo := systray.AddMenuItem("Check 1..2..", "Test Menu Item")

	mQuit := systray.AddMenuItem("Quit", "Quit bittray")

	go func() {
		for {
			select {
			case <-checkOneTwo.ClickedCh:
				fmt.Println("Clicked Check 1..2..")
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

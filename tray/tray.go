package tray

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/michaelsanford/bittray/polling"
	"github.com/michaelsanford/bittray/tray/assets/checkmark"
	"github.com/michaelsanford/bittray/tray/assets/star"
)

func Run() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(iconCheckmark.Data)
	systray.SetTitle("BitTray")
	systray.SetTooltip("Loading...")

	checkOneTwo := systray.AddMenuItem("Check 1..2..", "Test Menu Item")

	mQuit := systray.AddMenuItem("Quit", "Quit bittray")

	go func() {
		for s := range polling.Poll() {
			if len(s) > 0 {
				systray.SetTooltip(fmt.Sprintf("%d PRs waiting...", len(s)))
				systray.SetIcon(iconStar.Data)
			}

			for i := 0; i < len(s); i++ {
				//fmt.Println(s.label)
				//systray.AddMenuItem(s, "Pull Request")
			}
		}
	}()

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
	fmt.Println("Thank you; come again.")
}

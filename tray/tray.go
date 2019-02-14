package tray

import (
	"fmt"
	"github.com/gen2brain/dlgs"
	"github.com/michaelsanford/bittray/credentials"
	"github.com/michaelsanford/bittray/polling"
	"github.com/michaelsanford/bittray/tray/assets/checkmark"
	"github.com/michaelsanford/bittray/tray/assets/star"
	"github.com/michaelsanford/systray"
	"github.com/pkg/browser"
)

func Run() {
	systray.Run(onReady, onExit)
}

func onReady() {

	systray.SetIcon(checkmark.Icon)
	systray.SetTitle("Bittray")
	systray.SetTooltip("Loading...")

	mQuit := systray.AddMenuItem("Quit", "Quit Bittray")
	mReset := systray.AddMenuItem("Reset", "Reset Bittray to factory defaults")
	systray.AddSeparator()
	mStash := systray.AddMenuItem("Go to BitBucket", "Review your open Pull Requests")

	go func() {
		for prs := range polling.Poll() {
			if len(prs) > 0 {
				var plural string
				if len(prs) > 1 {
					plural = "s"
				}
				systray.SetTooltip(fmt.Sprintf("%d PR%s waiting...", len(prs), plural))
				systray.SetIcon(star.Icon)
			} else {
				systray.SetTooltip("Pull Request queue clear!")
				systray.SetIcon(checkmark.Icon)
			}
		}
	}()

	go func() {
		_, _, url := credentials.GetCred()
		for {
			select {
			case <-mStash.ClickedCh:
				err := browser.OpenURL(url)
				if err != nil {
					panic(err)
				}
			case <-mReset.ClickedCh:
				yes, err := dlgs.Question("Reset?", "Reset Bittray to factory defaults?", true)
				if err != nil {
					panic(err.Error())
				}
				if yes {
					credentials.DestroyCred()
					dlgs.Info("Reset", "Bittray has been reset and will now quit.")
					systray.Quit()
					return
				}
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

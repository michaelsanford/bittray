package tray

import (
	"fmt"
	"github.com/gen2brain/dlgs"
	"github.com/michaelsanford/bittray/credentials"
	"github.com/michaelsanford/bittray/polling"
	"github.com/michaelsanford/bittray/tray/assets"
	"github.com/michaelsanford/systray"
	"github.com/pkg/browser"
)

// Run starts the system tray and polling flow
func Run() {
	systray.Run(onReady, onExit)
}

func onReady() {

	systray.SetIcon(icon.Lock)
	systray.SetTitle("Bittray")
	systray.SetTooltip("Locked")

	mQuit := systray.AddMenuItem("Quit", "Quit Bittray")
	mReset := systray.AddMenuItem("Reset", "Reset Bittray to factory defaults")
	systray.AddSeparator()
	mStash := systray.AddMenuItem("Go to BitBucket", "Review your open Pull Requests")

	go func() {
		warned := false

		for count := range polling.Poll() {
			if count > 0 {
				warned = false
				var plural string
				if count > 1 {
					plural = "s"
				}
				systray.SetIcon(icon.Alarm)
				systray.SetTooltip(fmt.Sprintf("%d PR%s waiting...", count, plural))
			} else if count == 0 {
				warned = false
				systray.SetIcon(icon.Checkmark)
				systray.SetTooltip("Pull Request queue clear!")
			} else if count == -1 {
				if !warned {
					warned = true
					systray.SetIcon(icon.Lock)
					systray.SetTooltip("Locked")
					dlgs.Error("Bitbucket Error", "There was a problem contacting the API")
				}
			}
		}
	}()

	go func() {
		_, url := credentials.GetConfig()
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
					credentials.DestroyConfig()
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

package tray

import (
	"fmt"
	"github.com/gen2brain/dlgs"
	"github.com/getlantern/systray"
	"github.com/michaelsanford/bittray/config"
	"github.com/michaelsanford/bittray/polling"
	"github.com/michaelsanford/bittray/tray/assets"
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
	mUpdate := systray.AddMenuItem("Update Available...", "Get new version of Bittray")
	mUpdate.Hide()
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
				message := fmt.Sprintf("%d Pull Request%s", count, plural)

				systray.SetIcon(icon.Alarm)
				systray.SetTooltip(message)
				mStash.SetTitle("Review " + message)
				mStash.SetTooltip(message + " waiting...")
			} else if count == 0 {
				warned = false
				systray.SetIcon(icon.Checkmark)
				systray.SetTooltip("Pull Request queue clear!")
				mStash.SetTitle("Go to Bitbucket")
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
		_, url := config.GetConfig()
		for {
			select {
			case <-mStash.ClickedCh:
				err := browser.OpenURL(url)
				if err != nil {
					panic(err)
				}
			case <-mUpdate.ClickedCh:
				err := browser.OpenURL(config.DocsURL)
				if err != nil {
					panic(err)
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

	// Just check once on startup
	updateAvailable, latestTag := polling.CheckForUpdate()
	if updateAvailable {
		updateMsg := fmt.Sprintf("Update to %s available!", latestTag)
		mUpdate.SetTitle(updateMsg)
		mUpdate.Show()
	}
}

func onExit() {
	fmt.Println("Thank you; come again.")
}

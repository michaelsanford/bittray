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

		for bbResponse := range polling.Poll() {
			if bbResponse > 0 {
				warned = false

				var plural string
				if bbResponse > 1 {
					plural = "s"
				}
				message := fmt.Sprintf("%d Pull Request%s", bbResponse, plural)

				systray.SetIcon(icon.Alarm)
				systray.SetTooltip(message)
				mStash.SetTitle("Review " + message)
				mStash.SetTooltip(message + " waiting...")
			} else if bbResponse == 0 {
				warned = false
				systray.SetIcon(icon.Checkmark)
				systray.SetTooltip("Pull Request queue clear!")
				mStash.SetTitle("Go to Bitbucket")
			} else if bbResponse == -1 {
				if !warned {
					warned = true
					systray.SetIcon(icon.Lock)
					systray.SetTooltip("Locked")
					dlgs.Error("Bitbucket Error", "There was a problem contacting the API")
				}
			} else if bbResponse == -2 {
				systray.SetIcon(icon.Lock)
				systray.SetTooltip("Not Authorized")
				dlgs.Error("Not Authorized", "Wrong password. Quit and try again!")
				systray.Quit()
			} else if bbResponse == -3 {
				systray.SetIcon(icon.Rate)
				systray.SetTooltip("Rate Limited!")
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
	updateAvailable, latestTag, _ := polling.CheckForUpdate()
	if updateAvailable {
		updateMsg := fmt.Sprintf("Update to %s available!", latestTag)
		mUpdate.SetTitle(updateMsg)
		mUpdate.Show()
	}
}

func onExit() {
	fmt.Println("Thank you; come again.")
}

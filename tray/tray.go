package tray

import (
	"fmt"
	"github.com/gen2brain/dlgs"

	//"github.com/gen2brain/dlgs"
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

	mAuthored := systray.AddMenuItem("Go to BitBucket", "Your pull requests")
	mReviewing := systray.AddMenuItem("Go to BitBucket", "Your review tasks")

	go func() {
		//warned := false

		p := make(chan polling.DashboardSummary)
		polling.Poll(p)

		select {
		case d := <-p:
			fmt.Printf("READ: %+v\n", d)

			switch d.StatusCode {
			case 200:
				// Reviewing
				{
					if d.Reviewing > 0 {
						var plural string
						if d.Reviewing > 1 {
							plural = "s"
						}
						message := fmt.Sprintf("%d pull request%s", d.Reviewing, plural)

						systray.SetIcon(icon.Alarm)
						systray.SetTooltip(message)

						mReviewing.SetTitle("Review " + message)
						mReviewing.SetTooltip(message + " waiting...")

					} else {
						mReviewing.Hide()
					}
				}

				// Authored
				{
					if d.OpenTasks > 0 {
						var plural string
						if d.OpenTasks > 1 {
							plural = "s"
						}
						message := fmt.Sprintf("%d open task%s", d.OpenTasks, plural)

						systray.SetIcon(icon.Alarm)
						systray.SetTooltip(message)

						mAuthored.SetTitle("Review " + message)
						mAuthored.SetTooltip(message + " under review...")

						mAuthored.Show()
					} else {
						mAuthored.Hide()
					}
				}

				if d.Reviewing == 0 && d.OpenTasks == 0 {
					systray.SetIcon(icon.Checkmark)
					systray.SetTooltip("Pull Request queue clear!")
				}

				break
			case 401:
			case 403:
				// Unauthorized
				systray.SetIcon(icon.Lock)
				systray.SetTooltip("Not Authorized")
				dlgs.Error("Unauthorized", "Flushing configuration, restart.")
				config.DestroyConfig()
				systray.Quit()
				break
			case 404:
				break
			case 429:
				systray.SetIcon(icon.Rate)
				systray.SetTooltip("Rate Limited!")
				break
			case 504:
				systray.SetIcon(icon.Rate)
				systray.SetTooltip("Request timed out; retrying...")
			default:
				_, _ = dlgs.Info("Unexpected", string(d.StatusCode))
				break
			}
		}
	}()

	go func() {
		_, url := config.GetConfig()
		for {
			select {
			case <-mReviewing.ClickedCh:
				err := browser.OpenURL(url)
				if err != nil {
					panic(err)
				}
			case <-mAuthored.ClickedCh:
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

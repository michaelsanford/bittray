package tray

import (
	"fmt"
	"github.com/michaelsanford/bittray/credentials"
	"github.com/michaelsanford/bittray/polling"
	"github.com/michaelsanford/bittray/tray/assets/checkmark"
	"github.com/michaelsanford/bittray/tray/assets/star"
	"github.com/michaelsanford/systray"
	"github.com/pkg/browser"
	"strings"
)

func Run() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(iconCheckmark.Data)
	systray.SetTitle("BitTray")
	systray.SetTooltip("Loading...")

	mQuit := systray.AddMenuItem("Quit", "Quit Bittray")
	systray.AddSeparator()
	mStash := systray.AddMenuItem("Go to BitBucket", "Review your open Pull Requests")

	go func() {
		var sb strings.Builder

		for prs := range polling.Poll() {
			if len(prs) > 0 {

				var authors = make(map[string]struct{}, len(prs))
				var projects = make(map[string]struct{}, len(prs))

				for i := 0; i < len(prs); i++ {
					if _, exists := projects[prs[i].Project]; !exists {
						projects[prs[i].Project] = struct{}{}
					}

					if _, exists := authors[prs[i].Author]; !exists {
						projects[prs[i].Author] = struct{}{}
					}
				}

				sb.WriteString(fmt.Sprintf("%d PRs waiting in ", len(prs)))

				for project := range projects {
					sb.WriteString(project + " ")
				}

				for author := range authors {
					sb.WriteString(author + " ")
				}

				systray.SetTooltip(sb.String())
				systray.SetIcon(iconStar.Data)

				sb.Reset()
			} else {
				systray.SetTooltip("Pull Request queue clear!")
				systray.SetIcon(iconCheckmark.Data)
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

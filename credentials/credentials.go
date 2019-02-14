package credentials

import (
	"fmt"
	"github.com/danieljoos/wincred"
	"github.com/gen2brain/dlgs"
	"github.com/getlantern/systray"
	"net/url"
	"strings"
)

const credentialTarget string = "bt050"

type Auth struct {
	user string
	pass string
	url  string
}

func StoreCred(username string, password string, url string) {
	// TODO Replace this with AES (+ Registry?)
	cred := wincred.NewGenericCredential(credentialTarget)
	cred.UserName = strings.TrimSpace(username)
	cred.Comment = strings.TrimSpace(url)
	cred.CredentialBlob = []byte(strings.TrimSpace(password))
	err := cred.Write()

	if err != nil {
		panic(err)
	}

	return
}

func GetCred() (user string, pass string, url string) {
	cred, err := wincred.GetGenericCredential(credentialTarget)

	if err == nil {
		return cred.UserName, string(cred.CredentialBlob), cred.Comment
	}

	return "", "", ""
}

func AskCred() (ok bool) {

	dlgs.Warning("Bittray", "You're new here! Let's get you set up. You'll need to provide your Bitbucket username, password and URL.")

	username, ok, _ := askUser()
	if !ok {
		systray.Quit()
		return ok
	}

	password, ok, _ := askPass()
	if !ok {
		systray.Quit()
		return ok
	}

	address, ok, _ := askUrl()
	if !ok {
		systray.Quit()
		return ok
	}

	StoreCred(username, password, address)
	return ok
}

func askUser() (user string, ok bool, err error) {
	for user == "" {
		user, ok, err = dlgs.Entry("Username", "Please enter your BitBucket username", "")
		if user == "" && ok {
			dlgs.Error("Username missing", "Ok, so, without your username I can't log you in.\n\nTry again...")
		} else {
			return user, ok, err
		}
	}
	return user, ok, err
}

func askPass() (pass string, ok bool, err error) {
	for pass == "" {
		pass, ok, err = dlgs.Password("Password", "Please enter your Bitbucket password")
		if pass == "" && ok {
			dlgs.Error("Password missing", "You left the password field blank.\n\nThat's just...not going to work.")
		} else {
			return pass, ok, err
		}
	}
	return pass, ok, err
}

func askUrl() (pUrl string, ok bool, err error) {

	for pUrl == "" {
		pUrl, ok, err = dlgs.Entry("Bitbucket URL", "Enter your Bitbucket URL in exactly the format shown", "http://host.domain.com:7990")

		_, parsingErr := url.ParseRequestURI(pUrl)
		if parsingErr != nil && ok {
			dlgs.Error("Bad URL Format", "Sorry, the url you provide must be exactly of the format provided below, with a port and no trailing slash.\n\nPlease retry.")
			pUrl = ""
		} else {
			return pUrl, ok, err
		}
	}

	return pUrl, ok, err
}

func DestroyCred() {
	cred, err := wincred.GetGenericCredential(credentialTarget)
	if err != nil {
		fmt.Println(err)
		return
	}

	cred.Delete()
}

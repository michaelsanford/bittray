package config

import (
	"fmt"
	"github.com/danieljoos/wincred"
	"github.com/gen2brain/dlgs"
	"net/url"
	"strings"
)

const credentialTarget string = "bittray:conf"

// AskConfig interactively asks configuration information from the user
func AskConfig() (ok bool) {

	dlgs.Warning("Bittray", "You're new here! Let's get you set up. You'll need to provide your Bitbucket username, password and URL.")

	username, ok, _ := askUser()
	if !ok {
		return ok
	}

	address, ok, _ := askURL()
	if !ok {
		return ok
	}

	StoreConfig(username, address)
	return ok
}

// GetConfig retrieves the persisted configuration from the WCM
func GetConfig() (user string, url string) {
	cred, err := wincred.GetGenericCredential(credentialTarget)

	if err == nil {
		return cred.UserName, string(cred.CredentialBlob)
	}

	return "", ""
}

// StoreConfig persists the configuration to the WCM
func StoreConfig(username string, url string) {
	cred := wincred.NewGenericCredential(credentialTarget)
	cred.UserName = strings.TrimSpace(username)
	cred.CredentialBlob = []byte(strings.TrimSpace(url))
	err := cred.Write()

	if err != nil {
		panic(err)
	}

	return
}

func askUser() (user string, ok bool, err error) {
	for user == "" {
		user, ok, err = dlgs.Entry("Username", "Your BitBucket username:", "")
		if user == "" && ok {
			dlgs.Error("Username missing", "Ok, so, without your username I can't log you in.\n\nTry again...")
		} else {
			return user, ok, err
		}
	}
	return user, ok, err
}

func askURL() (pURL string, ok bool, err error) {

	for pURL == "" {
		pURL, ok, err = dlgs.Entry("Bitbucket URL", "Your Bitbucket URL:", "http://host.domain.com:7990")

		_, parsingErr := url.ParseRequestURI(pURL)
		if parsingErr != nil && ok {
			dlgs.Error("Bad URL Format", "Sorry, the url you provide must be exactly of the format provided below, with a port and no trailing slash.\n\nPlease retry.")
			pURL = ""
		} else {
			return pURL, ok, err
		}
	}

	return pURL, ok, err
}

// AskPass interactively asks the user for their Bitbucket password
func AskPass() (pass string, ok bool, err error) {
	for pass == "" {
		pass, ok, err = dlgs.Password("Bitbucket Password", "Your Bitbucket password:")
		if pass == "" || !ok {
			return "", ok, err
		}
	}
	return pass, ok, err
}

// DestroyConfig removes the persisted configuration in the WCM
func DestroyConfig() {
	cred, err := wincred.GetGenericCredential(credentialTarget)
	if err != nil {
		fmt.Println(err)
		return
	}

	cred.Delete()
}

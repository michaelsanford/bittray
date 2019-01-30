package credentials

import (
	"fmt"
	"github.com/danieljoos/wincred"
)

const credentialTarget string = "github.com/michaelsanford/bittray"

func StoreCred(url string, username string, password string) {
	cred := wincred.NewGenericCredential(credentialTarget)
	cred.UserName = username
	cred.Comment = url
	cred.CredentialBlob = []byte(password)
	err := cred.Write()

	if err != nil {
		panic(err)
	}
}

func GetCred() (auth map[string]string) {
	auth = make(map[string]string)

	cred, err := wincred.GetGenericCredential(credentialTarget)

	if err == nil {
		auth["user"] = cred.UserName
		auth["pass"] = string(cred.CredentialBlob)
		auth["url"] = cred.Comment
		return auth
	}

	return nil
}

func DestroyCred() {
	cred, err := wincred.GetGenericCredential(credentialTarget)
	if err != nil {
		fmt.Println(err)
		return
	}
	cred.Delete()
}

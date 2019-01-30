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

func GetCred() {
	cred, err := wincred.GetGenericCredential(credentialTarget)
	if err == nil {
		output := fmt.Sprintf("%s:%s@%s", cred.UserName, cred.CredentialBlob, cred.Comment)
		fmt.Println(output)
	}
}

func DestroyCred() {
	cred, err := wincred.GetGenericCredential(credentialTarget)
	if err != nil {
		fmt.Println(err)
		return
	}
	cred.Delete()
}

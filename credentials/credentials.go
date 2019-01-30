package main

import (
	"fmt"
	"github.com/danieljoos/wincred"
)

func StoreCred(url string, username string, password string) {
	cred := wincred.NewGenericCredential("bittray")
	cred.CredentialBlob = []byte(username + "|" + url + "|" + password)
	err := cred.Write()

	if err != nil {
		fmt.Println(err)
	}
}

func GetCred() {
	cred, err := wincred.GetGenericCredential("bittray")
	if err == nil {
		fmt.Println(string(cred.CredentialBlob))
	}
}

func DestroyCred() {
	cred, err := wincred.GetGenericCredential("bittray")
	if err != nil {
		fmt.Println(err)
		return
	}
	cred.Delete()
}
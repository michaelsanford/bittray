package credentials

import (
	"bufio"
	"fmt"
	"github.com/danieljoos/wincred"
	"github.com/michaelsanford/bittray/console"
	"net/url"
	"os"
	"strings"
)

const credentialTarget string = "github.com/michaelsanford/bittray"

type Auth struct {
	user string
	pass string
	url  string
}

func StoreCred(username string, password string, url string) {
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

func AskCred() {
	console.Show()

	fmt.Println("Looks like you're a new user; welcome!")
	fmt.Println("Let's get you set up...")
	fmt.Println("")
	fmt.Println("You will be asked to provide your BitBucket Enterprise (Stash) credentials.")
	fmt.Println("This information is stored directly in the Windows Credential Manager.")
	fmt.Println("For more information, see")
	fmt.Println("- https://support.microsoft.com/en-ca/help/4026814/windows-accessing-credential-manager ")
	fmt.Println("- https://github.com/michaelsanford/bittray ")
	fmt.Println("")

	reader := bufio.NewReader(os.Stdin)

	pUser, _ := askUser(reader)

	pPass, _ := askPass(reader)

	pUrl := askUrl(reader)
	for pUrl == "" {
		pUrl = askUrl(reader)
	}

	fmt.Println("If everything looks correct, hit [ENTER] to continue.")
	fmt.Println("(If you see an error, press CTRL-C and launch the app again.)")
	reader.ReadString('\n')

	StoreCred(pUser, pPass, pUrl)

	// TODO Test API Connection

	return
}

func askUser(reader *bufio.Reader) (user string, err error) {
	fmt.Print("Username: ")
	return reader.ReadString('\n')
}

func askPass(reader *bufio.Reader) (pass string, err error) {
	fmt.Print("Password: ")
	return reader.ReadString('\n')
}

func askUrl(reader *bufio.Reader) (vUrl string) {
	fmt.Print("URL [as http://host.server.net:7990/]: ")
	pUrl, _ := reader.ReadString('\n')

	u, err := url.ParseRequestURI(pUrl)
	if err != nil {
		fmt.Println(" Sorry, the URL must be EXACTLY of the format [http://host.domain.com:port/]. \n Please try again.")
		return ""
	}

	return fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, u.Path)
}

func DestroyCred() {
	cred, err := wincred.GetGenericCredential(credentialTarget)
	if err != nil {
		fmt.Println(err)
		return
	}

	cred.Delete()
}

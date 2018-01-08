package authenticate

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/okeyonyia123/cityrescue/models"

	"github.com/gorilla/sessions"
)

var SessionStore *sessions.FilesystemStore

//Here is where you create a session
func CreateUserSession(u *models.User, sessionID string, w http.ResponseWriter, r *http.Request) error {
	gfSession, err := SessionStore.Get(r, "cityrescue-session")

	if err != nil {
		log.Print(err)
	}

	gfSession.Values["sessionID"] = sessionID
	gfSession.Values["username"] = u.Username
	gfSession.Values["firstName"] = u.FirstName
	gfSession.Values["lastName"] = u.LastName
	gfSession.Values["email"] = u.Email

	err = gfSession.Save(r, w)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func init() {

	SessionStore = sessions.NewFilesystemStore("", []byte(os.Getenv("CITYRESCUE_HASH_KEY")))
	fmt.Println([]byte(os.Getenv("CITYRESCUE_HASH_KEY")))
	fmt.Println(os.Getenv("CITYRESCUE_HASH_KEY"))
}

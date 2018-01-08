package authenticate

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/okeyonyia123/cityrescue/models"

	"github.com/gorilla/securecookie"
	"github.com/subosito/gotenv"
)

var hashKey []byte
var blockKey []byte
var s *securecookie.SecureCookie

func CreateSecureCookie(u *models.User, sessionID string, w http.ResponseWriter, r *http.Request) error {

	value := map[string]string{
		"username": u.Username,
		"sid":      sessionID,
	}

	if encoded, err := s.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:     "session",
			Value:    encoded,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
	} else {
		log.Print(err)
		return err
	}

	return nil

}

func ReadSecureCookieValues(w http.ResponseWriter, r *http.Request) (map[string]string, error) {
	if cookie, err := r.Cookie("cityrescue-session"); err == nil {
		value := make(map[string]string)

		if err = s.Decode("cityrescue-session", cookie.Value, &value); err == nil {
			fmt.Println("Found Erroor!!!!!!!!!!!!!!!")
			return value, nil
		} else {
			return nil, err
		}
	} else {
		fmt.Printf("here!!!!!!!!!")
		return nil, nil
	}
}

func ExpireSecureCookie(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:   "cityrescue-session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", 301)
}

func init() {
	gotenv.Load()

	log.Println(os.Getenv("CITYRESCUE_HASH_KEY"))
	log.Println(os.Getenv("CITYRESCUE_BLOCK_KEY"))
	//For authenticating the Cookie Values
	hashKey = []byte(os.Getenv("CITYRESCUE_HASH_KEY"))
	//for encrypting the cookie values
	blockKey = []byte(os.Getenv("CITYRESCUE_BLOCK_KEY"))

	s = securecookie.New(hashKey, blockKey)
}

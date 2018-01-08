package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/okeyonyia123/cityrescue/shared/authenticate"
)

func GatedContentHandler(next http.HandlerFunc) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		shouldRedirectToLogin := false

		secureCookieMap, err := authenticate.ReadSecureCookieValues(w, r)
		if err != nil {
			log.Print(err)
		}

		//fmt.Printf("secure cookie contents: %+v\n", secureCookieMap)

		// Check if the sid key which is used to store the session id value
		// has been populated in the map using the comma ok idiom
		if _, ok := secureCookieMap["sid"]; ok == true {

			gfSession, err := authenticate.SessionStore.Get(r, "cityrescue-session")

			fmt.Printf("cityrescue session values: %+v\n", gfSession.Values)
			if err != nil {
				log.Print(err)
				return
			}

			// Check if the session id stored in the secure cookie matches
			// the id and username on the server-side session
			if gfSession.Values["sessionID"] == secureCookieMap["sid"] && gfSession.Values["username"] == secureCookieMap["username"] {
				next(w, r)
			} else {
				fmt.Printf("one")
				shouldRedirectToLogin = true
			}

		} else {
			fmt.Printf("Two")
			shouldRedirectToLogin = true

		}

		if shouldRedirectToLogin == true {
			http.Redirect(w, r, "/login", 302)
			fmt.Printf("Just Note That Cookie wasn't created successfully")
		}

	})

}

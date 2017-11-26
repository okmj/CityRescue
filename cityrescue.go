package main

import (
	"net/http"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/okeyonyia123/CityRescue/endpoints"
	"github.com/okeyonyia123/CityRescue/handlers"
	"github.com/okeyonyia123/CityRescue/middleware"
)

const (
	WEBSERVERPORT = ":8080"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	//Handle favicon request we have a favicon to work with
	http.Handle("/favicon.ico", http.NotFoundHandler())
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("GET,POST")
	r.HandleFunc("/signup", handlers.SignUpHandler).Methods("GET", "POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	//r.HandleFunc("/feed", handlers.FeedHandler).Methods("GET")
	//r.HandleFunc("/friends", handlers.FriendsHandler).Methods("GET")
	//r.HandleFunc("/find", handlers.FindHandler).Methods("GET,POST")
	r.HandleFunc("/profile", handlers.MyProfileHandler).Methods("GET")
	r.HandleFunc("/profile/{username}", handlers.ProfileHandler).Methods("GET")
	r.HandleFunc("/triggerpanic", handlers.TriggerPanicHandler).Methods("GET")
	r.HandleFunc("/foo", handlers.FooHandler).Methods("GET")

	//styling
	r.PathPrefix("/bootstrap/").Handler(http.StripPrefix("/bootstrap/", http.FileServer(http.Dir("./bootstrap"))))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("/restapi/disasterrecovery/{username}", endpoints.FetchPostsEndpoint).Methods("GET")
	r.HandleFunc("/restapi/disasterrecovery/{postid}", endpoints.CreatePostEndpoint).Methods("POST")
	r.HandleFunc("/restapi/disasterrecovery/{postid}", endpoints.UpdatePostEndpoint).Methods("PUT")
	r.HandleFunc("/restapi/disasterrecovery/{postid}", endpoints.DeletePostEndpoint).Methods("DELETE")

	//http.Handle("/", r)
	//http.Handle("/", ghandlers.LoggingHandler(os.Stdout, r))
	//http.Handle("/", middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r)))
	http.Handle("/", middleware.ContextExampleHandler(middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r))))

	http.ListenAndServe(WEBSERVERPORT, nil)

}

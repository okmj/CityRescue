package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/okeyonyia123/cityrescue/shared"
	"github.com/okeyonyia123/cityrescue/shared/datastore"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/okeyonyia123/cityrescue/endpoints"
	"github.com/okeyonyia123/cityrescue/handlers"
	"github.com/okeyonyia123/cityrescue/middleware"
)

const (
	WEBSERVERPORT = ":8080"
)

func main() {

	//Connect To DATABASE
	//db, err := datastore.NewDatastore(datastore.MYSQL, "cityrescue:cityrescue@/cityrescue")
	db, err := datastore.NewDatastore(datastore.MONGODB, "192.168.43.229:27017")
	fmt.Println(db)

	if err != nil {
		log.Print(err)
	}

	defer db.Close()

	//Set the Working DB to the Established Connection to db
	//This will be used to parse SignUpHandle(argument)  with the datastore connection as a dependency injection
	env := shared.Env{DB: db}
	fmt.Println(env)

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)

	//Handle favicon request we have a favicon to work with
	http.Handle("/favicon.ico", http.NotFoundHandler())

	r.Handle("/signup", handlers.SignUpHandler(&env)).Methods("GET", "POST")
	r.Handle("/login", handlers.LoginHandler(&env)).Methods("GET", "POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")

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

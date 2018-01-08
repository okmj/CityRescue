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
	WEBSERVERPORT = ":8084"
)

func main() {

	//Connect To DATABASE
	//db, err := datastore.NewDatastore(datastore.MYSQL, "cityrescue:cityrescue@/cityrescue")
	db, err := datastore.NewDatastore(datastore.MONGODB, "159.203.92.34:27017")
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
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET", "POST")

	//gated Resources for only logged in Users
	r.HandleFunc("/request/{username}", handlers.RequestHandler).Methods("GET", "POST")
	r.HandleFunc("/helper/{username}", handlers.HelperHandler).Methods("GET", "POST")
	r.Handle("/profile", middleware.GatedContentHandler(handlers.MyProfileHandler)).Methods("GET")
	r.Handle("/profile/{username}", middleware.GatedContentHandler(handlers.ProfileHandler)).Methods("GET")

	r.HandleFunc("/profile", handlers.MyProfileHandler).Methods("GET")
	r.HandleFunc("/profile/{username}", handlers.ProfileHandler).Methods("GET")
	r.HandleFunc("/triggerpanic", handlers.TriggerPanicHandler).Methods("GET")
	r.HandleFunc("/foo", handlers.FooHandler).Methods("GET")

	//styling
	r.PathPrefix("/bootstrap/").Handler(http.StripPrefix("/bootstrap/", http.FileServer(http.Dir("./bootstrap"))))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	//RESTAPI
	r.Handle("/restapi/disasterrecovery/request", endpoints.CreateRequestEndPoint(&env)).Methods("POST")
	r.Handle("/restapi/disasterrecovery/allrequests/{username}", endpoints.GetRequestEndPoint(&env)).Methods("GET")
	r.Handle("/restapi/disasterrecovery/pendingrequests/{{.username}}", endpoints.GetPendingRequestEndPoint(&env)).Methods("GET")
	r.HandleFunc("/restapi/disasterrecovery/{postid}", endpoints.UpdatePostEndpoint).Methods("PUT")
	r.HandleFunc("/restapi/disasterrecovery/{postid}", endpoints.DeletePostEndpoint).Methods("DELETE")

	// Clean Logging of server activities and performance
	http.Handle("/", middleware.ContextExampleHandler(middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r))))

	//Authentication middleware
	//loggedRouter := ghandlers.LoggingHandler(os.Stdout, r)
	//stdChain := alice.New(middleware.PanicRecoveryHandler)
	//http.Handle("/", stdChain.Then(loggedRouter))

	//Starting the server on port 8080
	http.ListenAndServe(WEBSERVERPORT, nil)

}

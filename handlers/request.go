package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var m = make(map[string]string)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	m["PageTitle"] = "Victim Page"
	vars := mux.Vars(r)
	username := vars["username"]
	fmt.Println(username)
	//w.Write([]byte(username))
	m["username"] = username
	fmt.Println(m)
	RenderTemplate(w, "./views/request.html", m)

}

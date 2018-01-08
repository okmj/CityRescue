package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var infos = make(map[string]string)

func HelperHandler(w http.ResponseWriter, r *http.Request) {
	infos["PageTitle"] = "Helper Page"
	vars := mux.Vars(r)
	username := vars["username"]
	fmt.Println(username)
	//w.Write([]byte(username))
	infos["username"] = username
	fmt.Println(infos)
	RenderTemplate(w, "./views/helper.html", infos)

}

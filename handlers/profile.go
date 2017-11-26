package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]
	fmt.Println(username)
	w.Write([]byte(username))

}

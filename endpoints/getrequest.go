package endpoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/okeyonyia123/cityrescue/handlers"
	"github.com/okeyonyia123/cityrescue/shared"
)

func PopulateGetRequestFormFields(resp http.ResponseWriter, req *http.Request, rF *RequestForm, env *shared.Env, username string) {
	//connect to database
	posts, err := env.DB.GetPost(username)

	if err != nil {
		//fmt.Println("Here where the error occured")
		log.Print(err)

	}

	fmt.Println(posts)

	handlers.RenderTemplate(resp, "./views/allrequests.html", posts)

}

func GetRequestEndPoint(env *shared.Env) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		rF := RequestForm{}
		rF.FieldNames = []string{"username", "category", "city", "address"}
		rF.Fields = make(map[string]string)

		vars := mux.Vars(req)
		username := vars["username"]
		fmt.Println(username)
		//w.Write([]byte(username))

		PopulateGetRequestFormFields(resp, req, &rF, env, username)

	})

}

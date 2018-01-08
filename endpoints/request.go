package endpoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/okeyonyia123/cityrescue/handlers"
	"github.com/okeyonyia123/cityrescue/models"
	"github.com/okeyonyia123/cityrescue/shared"
)

type RequestForm struct {
	FieldNames []string
	Fields     map[string]string
}

//Display Request Confirmation
func DisplayRequestConfirmation(w http.ResponseWriter, r *http.Request, rF *RequestForm) {
	handlers.RenderTemplate(w, "./views/requestconfirmation.html", rF)

}

func DisplayRequestForm(w http.ResponseWriter, r *http.Request, rf *RequestForm) {
	//tmpl.ExecuteTemplate(w, "signupform.html", s)
	handlers.RenderTemplate(w, "./views/request.html", rf)

}

func PopulateRequestFormFields(r *http.Request, rF *RequestForm) {

	for _, fieldName := range rF.FieldNames {
		rF.Fields[fieldName] = r.FormValue(fieldName)
	}

}

// ProcessSignUpForm
func ProcessRequestForm(w http.ResponseWriter, r *http.Request, rF *RequestForm, env *shared.Env) {

	p := models.NewPost(r.FormValue("username"), r.FormValue("category"), r.FormValue("city"), r.FormValue("address"))
	fmt.Println("post: ", p)

	//Just checking to make sure we still got a connection to the database
	fmt.Println(env)

	err := env.DB.CreatePost(p)

	if err != nil {
		//fmt.Println("Here where the error occured")
		log.Print(err)

	}

	// Display form confirmation message
	DisplayRequestConfirmation(w, r, rF)

}

func PostRequestForm(w http.ResponseWriter, r *http.Request, rF *RequestForm, e *shared.Env) {
	fmt.Println("Posting Request")
	//Populate the Reguest Form
	PopulateRequestFormFields(r, rF)
	ProcessRequestForm(w, r, rF, e)

}

func CreateRequestEndPoint(env *shared.Env) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		rF := RequestForm{}
		rF.FieldNames = []string{"username", "category", "city", "address"}
		rF.Fields = make(map[string]string)

		switch req.Method {

		case "GET":
			DisplayRequestForm(resp, req, &rF)
		case "POST":
			PostRequestForm(resp, req, &rF, env)
		default:
			DisplayRequestForm(resp, req, &rF)
		}

	})

}

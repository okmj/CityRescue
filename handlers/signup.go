package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/okeyonyia123/cityrescue/models"
	"github.com/okeyonyia123/cityrescue/shared"
	"github.com/okeyonyia123/cityrescue/validationkit"
)

type SignUpForm struct {
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

//var tmpl = template.Must(template.ParseFiles("./views"))

// DisplaySignUpForm displays the Sign Up form
func DisplaySignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	//tmpl.ExecuteTemplate(w, "signupform.html", s)
	RenderTemplate(w, "./views/signupform.html", s)

}

func DisplayConfirmation(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	RenderTemplate(w, "./views/signupconfirmation.html", s)
	//tmpl.ExecuteTemplate(w, "signupconfirmation.html", s)
}

func PopulateFormFields(r *http.Request, s *SignUpForm) {

	for _, fieldName := range s.FieldNames {
		s.Fields[fieldName] = r.FormValue(fieldName)
	}

}

// ValidateSignUpForm validates the Sign Up form's fields
func ValidateSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm, env *shared.Env) {

	PopulateFormFields(r, s)
	// Check if username was filled out
	if r.FormValue("username") == "" {
		s.Errors["usernameError"] = "The username field is required."
	}

	// Check if first name was filled out
	if r.FormValue("firstName") == "" {
		s.Errors["firstNameError"] = "The first name field is required."
	}

	// Check if last name was filled out
	if r.FormValue("lastName") == "" {
		s.Errors["lastNameError"] = "The last name field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("email") == "" {
		s.Errors["emailError"] = "The e-mail address field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("password") == "" {
		s.Errors["passwordError"] = "The password field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("confirmPassword") == "" {
		s.Errors["confirmPasswordError"] = "The confirm password field is required."
	}

	// Check username syntax
	if validationkit.CheckUsernameSyntax(r.FormValue("username")) == false {

		usernameErrorMessage := "The username entered has an improper syntax."
		if _, ok := s.Errors["usernameError"]; ok {
			s.Errors["usernameError"] += " " + usernameErrorMessage
		} else {
			s.Errors["usernameError"] = usernameErrorMessage
		}
	}

	// Check e-mail address syntax
	if validationkit.CheckEmailSyntax(r.FormValue("email")) == false {
		emailErrorMessage := "The e-mail address entered has an improper syntax."
		if _, ok := s.Errors["usernameError"]; ok {
			s.Errors["emailError"] += " " + emailErrorMessage
		} else {
			s.Errors["emailError"] = emailErrorMessage
		}
	}

	// Check if passord and confirm password field values match
	if r.FormValue("password") != r.FormValue("confirmPassword") {
		s.Errors["confirmPasswordError"] = "The password and confirm pasword fields do not match."
	}

	if len(s.Errors) > 0 {
		DisplaySignUpForm(w, r, s)
	} else {
		ProcessSignUpForm(w, r, s, env)
	}

}

// ProcessSignUpForm
func ProcessSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm, env *shared.Env) {

	// Now at this point, it means form submission is succefull so lets pass it into the DATABASE
	u := models.NewUser(r.FormValue("username"), r.FormValue("firstName"), r.FormValue("lastName"), r.FormValue("email"), r.FormValue("password"))
	fmt.Println("user: ", u)
	err := env.DB.CreateUser(u)

	if err != nil {
		log.Print(err)
	}

	user, err := env.DB.GetUser("kruti")
	if err != nil {
		log.Print(err)
	} else {
		fmt.Printf("Fetch User Result: %+v\n", user)
	}

	// Display form confirmation message
	DisplayConfirmation(w, r, s)

}

func SignUpHandler(env *shared.Env) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		s := SignUpForm{}
		s.FieldNames = []string{"username", "firstName", "lastName", "email"}
		s.Fields = make(map[string]string)
		s.Errors = make(map[string]string)

		switch req.Method {

		case "GET":
			DisplaySignUpForm(resp, req, &s)
		case "POST":
			ValidateSignUpForm(resp, req, &s, env)
		default:
			DisplaySignUpForm(resp, req, &s)
		}

	})

}

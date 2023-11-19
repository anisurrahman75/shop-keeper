package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/anisurrahman75/go-stock-management/api/data"
	"github.com/anisurrahman75/go-stock-management/api/models"
	"github.com/anisurrahman75/go-stock-management/pkg/auth"
	"html/template"
	"net/http"
)

func validateField(value string, fieldName string) ValidationResponse {
	if fieldName == "Email" || fieldName == "Password" {
		return ValidationResponse{
			Valid:   true,
			Message: fieldName + " is valid",
		}
	}
	return ValidationResponse{
		Valid:   true,
		Message: fieldName + " is invalid",
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	temp, err := template.ParseFiles("./api/templates/views/signup.html")
	if err != nil {
		panic(err)
	}
	if r.Method == http.MethodGet {
		if err != nil {
			panic(err)
		}
		err := temp.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	} else if r.Method == http.MethodPost {
		var formData map[string]string
		err := json.NewDecoder(r.Body).Decode(&formData)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		// Access and print the form data
		fmt.Println("Full Name:", formData["full_name"])
		fmt.Println("Email:", formData["email_add"])
		fmt.Println("Password:", formData["password"])

		validationResponses := make(map[string]ValidationResponse)
		validationResponses["full_name"] = validateField(formData["full_name"], "Full Name")
		validationResponses["email_add"] = validateField(formData["email_add"], "Email")
		validationResponses["password"] = validateField(formData["password"], "Password")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(validationResponses)
		if err != nil {
			panic(err)
		}
	}
}

func isEmailAndPasswordValid(email string, password string) bool {
	for _, user := range data.UserList {
		if user.Email == email && user.Password == password {
			return true
		}
	}
	return false
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./api/templates/views/signin.html")
	if err != nil {
		panic(err)
	}
	if r.Method == http.MethodGet {
		err = temp.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error on executing template", http.StatusBadRequest)
		}
	} else if r.Method == http.MethodPost {
		var formData map[string]string
		err := json.NewDecoder(r.Body).Decode(&formData)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		response := struct {
			AuthorizeUser bool `json:"authorize_user"`
		}{false}
		email := formData["email_add"]
		password := formData["password"]
		matchedUser := getMatchedUserWithEmailAndPass(email, password)

		if matchedUser != nil {
			response.AuthorizeUser = true
			if err := auth.GenerateJWTAndSetCookie(w, matchedUser); err != nil {
				http.Error(w, fmt.Sprintf("Failed to set JWT token in cookie. %s", err), http.StatusBadRequest)
				return
			}
		}
		fmt.Println("----res: ", response.AuthorizeUser)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			panic(err)
		}
	}
}

func getMatchedUserWithEmailAndPass(email string, pass string) *models.User {
	for _, userData := range data.UserList {
		if userData.Email == email && userData.Password == pass {
			return &userData
		}
	}
	return nil
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	auth.GetUserFromXCookieJWT(w, r)
	http.Redirect(w, r, "/signin", http.StatusSeeOther)
}

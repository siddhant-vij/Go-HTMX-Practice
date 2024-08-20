package main

import (
	"Auth_Validations/templates"
	"crypto/rand"
	"log"
	"math/big"
	"net/http"
	"strings"
)

func isEmailValid(email string) bool {
	return strings.Contains(email, "@")
}

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageComponent := templates.Page()
		pageComponent.Render(r.Context(), w)
	})

	http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")
		if email != "" && !isEmailValid(email) {
			emailError := templates.EmailError()
			emailError.Render(r.Context(), w)
			return
		} else if password != "" && len(password) < 8 {
			passError := templates.PasswordError()
			passError.Render(r.Context(), w)
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")

		var errors []string = make([]string, 0)
		if !isEmailValid(email) {
			errors = append(errors, "Email is not valid")
		}
		if len(password) < 8 {
			errors = append(errors, "Password is too short")
		}

		if len(errors) > 0 {
			loginErrors := templates.LoginErrors(errors)
			loginErrors.Render(r.Context(), w)
		}

		// Simulating a Server Error (DB/Network Error)
		randomNumBI, _ := rand.Int(rand.Reader, big.NewInt(2))
		randomNum := randomNumBI.Int64()
		if randomNum == 0 {
			w.Header().Set("HX-Retarget", ".control")
			w.Header().Set("HX-Reswap", "beforebegin")
			serverError := templates.ServerError()
			serverError.Render(r.Context(), w)
		}

		w.Header().Set("HX-Redirect", "/authenticated")
	})

	http.HandleFunc("/authenticated", func(w http.ResponseWriter, r *http.Request) {
		pageComponent := templates.Authenticated()
		pageComponent.Render(r.Context(), w)
	})

	log.Println("Starting server on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

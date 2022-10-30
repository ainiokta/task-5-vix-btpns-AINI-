package controllers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/ainiokta/task-5-vix-btpns-AINI-/app"
	"github.com/ainiokta/task-5-vix-btpns-AINI-/database"
	"github.com/ainiokta/task-5-vix-btpns-AINI-/helper"
	"github.com/ainiokta/task-5-vix-btpns-AINI-/models"
	"golang.org/x/crypto/bcrypt"
)

type Userinput struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

var Usermodel = models.NewUsermodel()
var Validation = helper.NewValidation()

func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := database.Store.Get(r, database.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

		temp, _ := template.ParseFiles("pages/dashboard.html")
		temp.Execute(w, nil)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("pages/login.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		// proses login
		r.ParseForm()
		UserInput := &Userinput{
			Email:    r.Form.Get("email"),
			Password: r.Form.Get("password"),
		}

		var users app.Users
		Usermodel.Where(&users, "email", UserInput.Email)

		var message error
		if users.Username == "" {
			message = errors.New("Email atau Password salah!")
		} else {
			// pengecekan password
			errPassword := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(UserInput.Password))
			if errPassword != nil {
				message = errors.New("Email atau Password salah!")
			}
		}

		if message != nil {

			data := map[string]interface{}{
				"error": message,
			}

			temp, _ := template.ParseFiles("pages/login.html")
			temp.Execute(w, data)
		} else {
			// set session
			session, _ := database.Store.Get(r, database.SESSION_ID)

			session.Values["loggedIn"] = true
			session.Values["email"] = users.Email
			session.Values["username"] = users.Username
			session.Values["nama"] = users.Nama

			session.Save(r, w)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}

}

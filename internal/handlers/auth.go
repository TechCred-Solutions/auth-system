package handlers

import (
	"log"
	"net/http"

	"auth-system/internal/db"
	"auth-system/internal/models"
	"auth-system/internal/session"
)

// Register handles user registration
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Retrieve username and password from the form
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Create a new user instance with the provided username and password
		user := models.User{Username: username, Password: password}
		// Hash the user's password
		if err := user.HashPassword(); err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		// Save the user to the database
		db.DB.Create(&user)
		// Redirect to the login page upon successful registration
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// Login handles user login
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Retrieve username and password from the form
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Find the user in the database by username
		var user models.User
		db.DB.Where("username = ?", username).First(&user)

		if user.ID == 0 || !user.CheckPassword(password) {
			// If the user ID is 0 or the password check fails, return an unauthorized error
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Retrieve the session
		sess, err := session.GetSession(r)
		if err != nil {
			log.Println("Unable to Get Session")
			http.Error(w, "Please Wait ...", http.StatusNotFound)
			return
		}

		// Store the user ID in the session
		sess.Values["user_id"] = user.ID
		// Save the session
		sess.Save(r, w)
		// Redirect to the dashboard upon successful login
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

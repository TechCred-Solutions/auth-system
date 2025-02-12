package handlers

import (
	"net/http"

	"auth-system/internal/db"
	"auth-system/internal/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user := models.User{ Username: username, Password: password}
		if err := user.HashPassword(); err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		db.DB.Create(&user)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

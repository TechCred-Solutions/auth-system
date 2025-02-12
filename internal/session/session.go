package session

import (
    "net/http"

    "github.com/gorilla/sessions"
)

// Initialize a new cookie store with a secret key
var store = sessions.NewCookieStore([]byte("super-secret-key"))

// GetSession retrieves the session from the request
func GetSession(r *http.Request) (*sessions.Session, error) {
    // Get the session named "session-1" from the store
    session, err := store.Get(r, "session-1")
    if err != nil {
        return nil, err
    }

    return session, nil
}

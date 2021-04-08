package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// Adds CSRF protection to every POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// Loads and saves a session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave((next))
}

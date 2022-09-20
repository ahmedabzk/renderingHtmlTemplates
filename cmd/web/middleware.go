package main

import (
	
	"net/http"

	"github.com/justinas/nosurf"
)

// Nosurf adds csrf protection to all post requests
func Nosurf(next http.Handler) http.Handler{
	csrfhandler := nosurf.New(next)

	csrfhandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfhandler
}

func LoadSession(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}
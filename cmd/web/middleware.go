package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteToConsole logs the request data to the console
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestData := fmt.Sprintf(
			`
			Receiving a request with data...

			Request HTTP Method => %s 
			Request URI => %s 
			Request Headers => %s
			`,
			r.Method,
			r.RequestURI,
			r.Header,
		)

		log.Println(requestData)
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

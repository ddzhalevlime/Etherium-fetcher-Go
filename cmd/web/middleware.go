package main

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/golang-jwt/jwt"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, r, fmt.Errorf("%s", err))
				app.logger.Error("panic recovered", "stack", debug.Stack())
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.logger.Info("received request",
			"method", r.Method,
			"uri", r.URL.RequestURI(),
			"remote_addr", r.RemoteAddr,
		)

		next.ServeHTTP(w, r)
	})
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error(err.Error())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) validateToken(w http.ResponseWriter, r *http.Request) (string, error) {
	tokenString := r.Header.Get("AUTH_TOKEN")
	if tokenString == "" {
		return "", fmt.Errorf("no AUTH_TOKEN provided")
	}

	token, err := app.parseToken(tokenString)
	if err != nil {
		return "", fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("username claim not found or not a string")
	}

	return username, nil
}

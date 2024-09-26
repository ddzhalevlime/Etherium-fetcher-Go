package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /lime/eth", app.getEth)
	mux.HandleFunc("GET /lime/eth/{rlphex}", app.getEthRlp)
	mux.HandleFunc("GET /lime/all", app.getAll)

	mux.HandleFunc("POST /lime/authenticate", app.postAuth)
	mux.HandleFunc("GET /lime/my", app.getMy)
	mux.HandleFunc("POST /lime/savePerson", app.postSavePerson)
	mux.HandleFunc("GET /lime/listPersons", app.getPersonList)

	return app.recoverPanic(app.logRequest(mux))
}

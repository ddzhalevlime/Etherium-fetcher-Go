package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/golang-jwt/jwt"
)

func (app *application) getEth(w http.ResponseWriter, r *http.Request) {
	hashStrings, err := handleTransactionHashesQueryString(r)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	username, _ := app.validateToken(w, r)
	transactions, err := app.fetchTransactions(r.Context(), hashStrings)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	if username != "" {
		err = app.users.InsertTransactionIds(username, extractTransactionIds(transactions))
		if err != nil {
			app.serverError(w, r, err)
			return
		}
	}

	app.responseJSON(w, r, map[string]interface{}{"transactions": transactions})
}

func (app *application) getEthRlp(w http.ResponseWriter, r *http.Request) {
	rlpHex := r.PathValue("rlphex")
	rlpBytes, err := hex.DecodeString(rlpHex)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var hashStrings []string
	err = rlp.DecodeBytes(rlpBytes, &hashStrings)
	if err != nil {
		app.serverError(w, r, fmt.Errorf("failed to decode RLP data: %w", err))
		return
	}

	username, _ := app.validateToken(w, r)
	transactions, err := app.fetchTransactions(r.Context(), hashStrings)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	if username != "" {
		err = app.users.InsertTransactionIds(username, extractTransactionIds(transactions))
		if err != nil {
			app.serverError(w, r, err)
			return
		}
	}

	app.responseJSON(w, r, map[string]interface{}{"transactions": transactions})
}

func (app *application) getAll(w http.ResponseWriter, r *http.Request) {
	transactions, err := app.transactions.GetAll()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]interface{}{"transactions": transactions})
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) postAuth(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	dbUser, err := app.users.Get(creds.Username)
	if err != nil {
		app.clientError(w, http.StatusNotFound)
		return
	}

	if dbUser == nil || dbUser.PasswordHash != hashWithJwtSecret(creds.Password) {
		app.clientError(w, http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": dbUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(app.jwtSecret))
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func (app *application) getMy(w http.ResponseWriter, r *http.Request) {
	username, err := app.validateToken(w, r)
	if err != nil {
		app.clientError(w, http.StatusUnauthorized)
		return
	}

	user, err := app.users.Get(username)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	transactions, _ := app.transactions.GetMultipleByIDs(user.SearchedTransactionIds)

	app.responseJSON(w, r, map[string]interface{}{"transactions": transactions})
}

func (app *application) postSavePerson(w http.ResponseWriter, r *http.Request) {
	var person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	txHash, txStatus, err := app.contractInteractor.SetPersonInfo(person.Name, person.Age)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"txHash":   txHash,
		"txStatus": txStatus,
	})
}

func (app *application) getPersonList(w http.ResponseWriter, r *http.Request) {
	persons, err := app.personInfoEvents.GetAll()

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]interface{}{"persons": persons})
	if err != nil {
		app.serverError(w, r, err)
	}
}

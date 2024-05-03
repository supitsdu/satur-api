package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/supitsdu/satur-api/internal/model"
	"github.com/supitsdu/satur-api/internal/repository"
	"github.com/supitsdu/satur-api/internal/response"
	"go.mongodb.org/mongo-driver/mongo"
)

type HandleAccounts struct {
	repo repository.Actions
}

// Creates a new instance of HandleAccounts
func CreateHandleAccounts(repo repository.Actions) *HandleAccounts {
	return &HandleAccounts{repo: repo}
}

// GetMethod handles GET requests for accounts information
func (h *HandleAccounts) GetMethod(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	account, err := h.repo.GetAccount(username)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			response.WriteError(w, http.StatusNotFound, "Account not found")
		} else {
			response.WriteError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	response.WriteJSON(w, http.StatusOK, account)
}

// PostMethod handles POST requests to create a new account
func (h *HandleAccounts) PostMethod(w http.ResponseWriter, r *http.Request) {
	var account model.AccountPersonalData
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.repo.CreateAccount(&account); err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed to create account")
		return
	}

	response.WriteJSON(w, http.StatusCreated, "Finished action successfully")
}

// DeleteMethod handles DELETE requests to delete an account
func (h *HandleAccounts) DeleteMethod(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	if err := h.repo.DeleteAccount(username); err != nil {
		if err == mongo.ErrNoDocuments {
			response.WriteError(w, http.StatusNotFound, "Account not found")
			return
		}
		response.WriteError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	response.WriteJSON(w, http.StatusNoContent, nil)
}

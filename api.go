package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiServer struct {
	listenAddr	string
	store 		storage
}

func newApiServer(listenAddr string, store storage) *apiServer  {
	return &apiServer{
		listenAddr: listenAddr,
		store:		store,
	}
}

func (s *apiServer) run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccountByID))


	http.ListenAndServe(s.listenAddr, router)

	log.Println("JSON API server running on port ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *apiServer)  handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *apiServer)  handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.getAccounts()
	
	if err != nil {
		return nil
	}

	return writeJSON(w, http.StatusOK, accounts)
}

func (s *apiServer)  handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	fmt.Println(id)

	return writeJSON(w, http.StatusOK, &account{})
}

func (s *apiServer)  handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountRequest := new(createAccountRequest)

	if err := json.NewDecoder(r.Body).Decode(createAccountRequest); err != nil {
		return err
	}

	account := newAccount(createAccountRequest.FirstName, createAccountRequest.LastName)

	if err := s.store.createAccount(account); err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, account)
}

func (s *apiServer)  handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *apiServer)  handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func (http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string
}

func makeHTTPHandleFunc (f apiFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}
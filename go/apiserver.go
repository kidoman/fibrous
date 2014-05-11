package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type ApiServer struct {
	db *DB
}

func newApiServer(db *DB) *ApiServer {
	return &ApiServer{db}
}

func (a *ApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "POST" && r.URL.Path == "/users":
		a.createUser(w, r)
	case r.Method == "GET" && strings.HasPrefix(r.URL.Path, "/users/"):
		a.getUser(w, r)
	}
}

func (a *ApiServer) createUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.db.SaveUser(&u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "OK")
}

func (a *ApiServer) getUser(w http.ResponseWriter, r *http.Request) {
	str := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := a.db.LoadUser(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if err := json.NewEncoder(w).Encode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *ApiServer) Close() {
	a.db.Close()
}

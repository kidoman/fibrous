package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	pool := newPool(":6379")
	defer pool.Close()

	db := newDB(pool)
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		str := vars["id"]

		id, err := strconv.Atoi(str)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		u, err := db.LoadUser(id)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		if err := json.NewEncoder(w).Encode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := db.SaveUser(&u); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "OK")
	})

	http.ListenAndServe(":3000", r)
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocraft/web"
	"net/http"
	"runtime"
	"strconv"
)

type Context struct {
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	pool := newPool(":6379")
	defer pool.Close()

	db := newDB(pool)
	defer db.Close()

	r := web.New(Context{})

	r.Get("/users/:id", func(rw web.ResponseWriter, r *web.Request) {
		str := r.PathParams["id"]

		id, err := strconv.Atoi(str)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		u, err := db.LoadUser(id)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(rw).Encode(&u); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Post("/users", func(rw web.ResponseWriter, r *web.Request) {
		defer r.Body.Close()

		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if err := db.SaveUser(&u); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		fmt.Fprint(rw, "OK")
	})

	http.ListenAndServe(":3000", r)
}

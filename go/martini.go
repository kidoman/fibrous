package main

import (
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/encoder"
)

func main() {
	m := martini.Classic()

	pool := newPool(":6379")
	defer pool.Close()

	db := newDB(pool)
	defer db.Close()

	m.Use(func(c martini.Context, w http.ResponseWriter) {
		c.MapTo(encoder.JsonEncoder{}, (*encoder.Encoder)(nil))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	m.Map(db)

	m.Post("/users", binding.Bind(User{}), func(db *DB, u User) (int, string) {
		if err := db.SaveUser(&u); err != nil {
			return http.StatusInternalServerError, err.Error()
		}
		return http.StatusCreated, "OK"
	})

	m.Get("/users/:id", func(db *DB, params martini.Params, enc encoder.Encoder) (int, []byte) {
		str := params["id"]
		id, err := strconv.Atoi(str)
		if err != nil {
			return http.StatusBadRequest, []byte{}
		}

		u, err := db.LoadUser(id)
		if err != nil {
			return http.StatusNotFound, []byte{}
		}
		return http.StatusOK, encoder.Must(enc.Encode(u))
	})

	m.Run()
}

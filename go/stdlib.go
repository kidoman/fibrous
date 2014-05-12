package main

import (
	"runtime"

	"net/http"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	pool := newPool(":6379")
	defer pool.Close()

	db := newDB(pool)
	defer db.Close()

	apiServer := newApiServer(db)
	defer apiServer.Close()

	http.Handle("/", apiServer)
	http.ListenAndServe(":3000", nil)
}

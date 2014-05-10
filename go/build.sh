#!/bin/bash --login
set -e

go get github.com/garyburd/redigo/redis
go get github.com/youtube/vitess/go/pools
go get github.com/go-martini/martini
go get github.com/martini-contrib/binding
go get github.com/martini-contrib/encoder

go build -o stdlib db.go redis.go user.go apiserver.go stdlib.go
go build -o martini db.go redis.go user.go martini.go

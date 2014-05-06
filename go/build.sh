#!/bin/bash
set -e

go build -o stdlib db.go redis.go user.go apiserver.go stdlib.go
go build -o martini db.go redis.go user.go martini.go

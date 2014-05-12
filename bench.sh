#!/bin/bash
set -e

go get github.com/rakyll/boom

curl -XPOST -H "Content-Type: application/json" -d '{"id": 1, "name": "Rob Pike"}' http://127.0.0.1:3000/users

GOMAXPROCS=4 boom -n ${T:-100000} -c ${C:-50} http://127.0.0.1:3000/users/1

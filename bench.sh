#!/bin/bash
set -e

go get github.com/rakyll/boom

curl -XPOST -H "Content-Type: application/json" -d '{"id": 1, "name": "Rob Pike"}' http://127.0.0.1:3000/users

boom -n 100000 -c 100 http://127.0.0.1:3000/users/1

#!/bin/bash
set -e

go get github.com/rakyll/boom

curl -i -XPOST -H "Content-Type: application/json" -d '{"id": 1, "name": "Rob Pike"}' http://localhost:3000/users

boom -n 100000 -c 100 http://localhost:3000/users/1

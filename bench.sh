#!/bin/bash
set -e

curl -XPOST -H "Content-Type: application/json" -d '{"id": 1, "name": "Rob Pike"}' http://127.0.0.1:3000/users

wrk -c${C:-100} -d${D:-20s} --latency http://127.0.0.1:3000/users/1

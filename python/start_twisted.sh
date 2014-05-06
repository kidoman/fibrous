#!/bin/bash --login
set -e

./build.sh

. .env/bin/activate
python twisted_server.py

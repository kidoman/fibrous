#!/bin/bash --login
set -e

./build.sh

. ./.env/bin/activate
python tornado_server.py

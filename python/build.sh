#!/bin/bash --login
set -e

virtualenv .env
. ./.env/bin/activate
pip install -r requirements.txt

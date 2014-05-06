#!/bin/bash --login
set -e

sudo easy_install virtualenv
virtualenv .env
. .env/bin/activate
pip install -r requirements.txt

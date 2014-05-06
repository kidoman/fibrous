#!/bin/bash --login
set -e

./build.sh

nvm install v0.10.28
nvm use v0.10.28
node fiber

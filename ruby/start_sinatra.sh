#!/bin/bash --login
set -e

./build.sh

bundle exec ruby sinatra.rb

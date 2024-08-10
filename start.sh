#!/bin/sh

set -e # ensure that the script will exit immediately if a command return status 0

echo "start the app"
exec "$@" # take all parameters paste to the script and run it

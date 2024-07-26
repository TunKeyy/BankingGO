#!/bin/sh

set -e # ensure that the script will exit immediately if a command return status 0
source /app/app.env

# echo "run db migration"
# /app/migrate -path /app/migration -database "DB_SOURCE" -verbose up

echo "start the app"
exec "$@" # take all parameters paste to the script and run it

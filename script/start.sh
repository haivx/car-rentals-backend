#!/bin/sh
# instruction to make sure that the script will exit immediately  if any command return none zer0 status
set -e

echo "run db migration"

/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up


echo "start app"

# take all the parameter pass to the script and run it
exec "$@"


# dont forget set chmod +x start.sh to make this file excecutable
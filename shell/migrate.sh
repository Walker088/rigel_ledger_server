#!/usr/bin/env bash

set -a
[ -f .project.env ] && . .project.env
set +a

flywayBaseline() {
    docker run --rm --network host \
        -v "$(pwd)/flyway:/flyway/sql" \
        flyway/flyway:8.5.0-alpine baseline -url=jdbc:postgresql://$DB_HOST:$DB_PORT/$DB_NAME?ApplicationName=rigel_flyway -table=$DB_SCHEMA -user=$DB_USER -password=$DB_PW -locations=filesystem:/flyway/sql
}

flywayMigrate() {
    docker run --rm --network host \
        -v "$(pwd)/flyway:/flyway/sql" \
        flyway/flyway:8.5.0-alpine migrate -url=jdbc:postgresql://$DB_HOST:$DB_PORT/$DB_NAME?ApplicationName=rigel_flyway -table=$DB_SCHEMA -user=$DB_USER -password=$DB_PW -locations=filesystem:/flyway/sql
}

flywayInfo() {
    docker run --rm --network host \
        -v "$(pwd)/flyway:/flyway/sql" \
        flyway/flyway:8.5.0-alpine info -url=jdbc:postgresql://$DB_HOST:$DB_PORT/$DB_NAME?ApplicationName=rigel_flyway -table=$DB_SCHEMA -user=$DB_USER -password=$DB_PW -locations=filesystem:/flyway/sql
}

if [[ $# -eq 0 ]] ; then
    echo 'Please provide one of the arguments (e.g., bash shell/migrate.sh info):
    1 > info
    2 > baseline
    3 > migrate'

elif [[ $1 == info ]]; then
    flywayInfo

elif [[ $1 == baseline ]]; then
    flywayBaseline

elif [[ $1 == migrate ]]; then
    flywayMigrate

fi

# Clean the env variables
unset $(grep -v '^#' .project.env | sed -E 's/(.*)=.*/\1/' | xargs)

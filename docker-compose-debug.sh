#!/bin/sh

json="$(cat ./configs/firebase-service-dev.json)"

export FIREBASE_SERVICE="$json"
export DOCKER_RUN_COMMAND="realize start --run"
export DIR_ENV=debug

cp ./configs/.env.local.sample .env

docker-compose up


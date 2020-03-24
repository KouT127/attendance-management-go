#!/bin/sh

json="$(cat ./configs/firebase-service-dev.json)"

export FIREBASE_SERVICE="$json"

cp ./configs/.env.local.sample .env

docker-compose up


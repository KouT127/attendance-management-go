#!/bin/sh

json="$(cat ./configs/firebase-service-dev.json)"

export FIREBASE_SERVICE="$json"
export DIR_ENV="production"

cp ./configs/.env.local.sample .env

docker-compose up


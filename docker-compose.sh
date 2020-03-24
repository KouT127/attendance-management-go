#!/bin/sh

json="$(cat ./configs/firebase-service-dev.json)"

export FIREBASE_SERVICE="$json"

docker-compose up

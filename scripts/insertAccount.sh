#!/bin/bash
curl --location --request POST 'localhost:8090/v1/accounts' \
--header 'Content-Type: application/json' \
--data-raw '{
    "DocumentNumber" : "101010"
}'
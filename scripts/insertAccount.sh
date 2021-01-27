#!/bin/bash
curl --location --request POST 'localhost:8080/v1/accounts' \
--header 'Content-Type: application/json' \
--data-raw '{
    "DocumentNumber" : "10101010"
}'
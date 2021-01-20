#!/bin/bash
curl --location --request POST 'localhost:8080/v1/transaction' \
--header 'Content-Type: application/json' \
--data-raw '{
    "AccountID" : 4,
    "OperationsTypeID" : 1,
    "Amount" : 20.36
}'
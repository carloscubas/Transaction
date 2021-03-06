# Transaction

## Objectives

This project is a small example of a REST API using GoLang. Inside scripts folder there are some examples:

- create new account (scripts/insertAccount.sh)
- create new transaction (scripts/insertTransaction.sh)
- get exist account (scripts/getAccount.sh)
- get all operation types (scripts/getOperationTypes.sh)


## Run 
Before Run app open configs/dev.yaml file and rewrite de data.
### Local environment

        go run cmd/app/main.go
        
### Docker environment

        docker-compose -f docker/docker-compose.dev.yaml up

## Tests

### Unit test

        go test -v -cover $$(go list ./... | grep -v /vendor/)

### Docker Unit tests

        docker-compose -f docker/docker-compose.test.yaml run test
        
### References
    - https://github.com/sbecker/gin-api-demo
    - https://www.agiratech.com/blog/building-restful-api-service-golang-using-gin-gonic
    - https://github.com/gin-gonic/gin#installation
    - https://gorm.io/
    - https://medium.com/@cgrant/developing-a-simple-crud-api-with-go-gin-and-gorm-df87d98e6ed1
    - https://github.com/golang-migrate/migrate
    - https://semaphoreci.com/community/tutorials/test-driven-development-of-go-web-applications-with-gin
    - https://kpat.io/2019/06/testing-with-gin/
    - https://github.com/1000kit/docker-h2
    - https://golang.org/pkg/testing/
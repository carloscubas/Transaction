package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"transaction/internal/account"
	"transaction/internal/config"
	"transaction/test"
)

var database string
var connection string

func init (){
	sc, _ := config.LoadServiceConfig("../../../configs/dev.yaml")
	database = sc.Db.Database
	connection = sc.Db.Connection
}

func TestInsertAccount(t *testing.T) {
	setup := test.NewDbTestConfig(database, connection)
	ts := httptest.NewServer(setupServer(setup.Conn))

	setup.Before()

	values := map[string]string{"DocumentNumber": "654321"}
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(fmt.Sprintf("%s/v1/accounts", ts.URL), "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	var ac account.Account
	err = json.NewDecoder(resp.Body).Decode(&ac)
	if err != nil {
		t.Fatalf("Expected account Body")
	}

	if ac.DocumentNumber != "654321" {
		t.Fatalf("Expected expected, got %s but %s", "654321", ac.DocumentNumber)
	}

	setup.After()
}

func TestGetAccount(t *testing.T) {
	setup := test.NewDbTestConfig(database, connection)
	ts := httptest.NewServer(setupServer(setup.Conn))

	setup.Before()

	resp, err := http.Get(fmt.Sprintf("%s/v1/accounts/1", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	var ac account.Account
	err = json.NewDecoder(resp.Body).Decode(&ac)
	if err != nil {
		fmt.Printf(err.Error())
		t.Fatalf("Expected account Body")
	}

	if ac.DocumentNumber != "123456" {
		t.Fatalf("Expected expected, got %s but %s", "123456", ac.DocumentNumber)
	}
	setup.After()

}

func TestInsertTransaction(t *testing.T) {
	setup := test.NewDbTestConfig(database, connection)
	ts := httptest.NewServer(setupServer(setup.Conn))

	setup.Before()

	values := map[string]interface{}{"AccountID": 1, "OperationsTypeID": 1, "Amount": 20.36}
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(fmt.Sprintf("%s/v1/transaction", ts.URL), "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	var tr account.Transaction
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		t.Fatalf("Expected account Body")
	}

	if tr.Amount != -20.36 {
		t.Fatalf("Expected expected, got %f but %f", -20.36, tr.Amount)
	}
	setup.After()
}

func TestGetType(t *testing.T) {

	setup := test.NewDbTestConfig(database, connection)
	ts := httptest.NewServer(setupServer(setup.Conn))

	setup.Before()
	resp, err := http.Get(fmt.Sprintf("%s/v1/operationtypes", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	types := make([]account.OperationType, 0)
	err = json.NewDecoder(resp.Body).Decode(&types)
	if err != nil {
		t.Fatalf("Expected account Body")
	}
	if  len(types) != 4{
		t.Errorf("expected %d, got %d", 4, len(types))
	}
	setup.After()
}

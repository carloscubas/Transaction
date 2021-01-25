package config

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestLoadServiceConfig ...
func TestConfig(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		t.Error("Cannot create temporary file", err)
	}
	defer os.Remove(tmpFile.Name())

	// yaml test
	input := []byte(`service_name: "transaction"

server:
    address: ":8080"
    mode: "debug"

db:
    database: "mysql"
    connection: "user:password@tcp(127.0.0.1:3306)/db"`)

	if _, err = tmpFile.Write(input); err != nil {
		t.Error("Failed to write to temporary file", err)
	}

	svc, err := LoadServiceConfig(tmpFile.Name())
	if err != nil {
		t.Errorf("Wrong file")
	}

	//validate service name
	if len(svc.ServiceName) == 0 {
		t.Errorf("Wrong file")
	}
}

// TestLoadServiceConfigBadPath ...
func TestConfigBadPath(t *testing.T) {

	_, err := LoadServiceConfig("CDE")
	if err != nil {
		t.Log("Wrong file")
	}

}

// TestLoadServiceConfigBadBody ...
func TestConfigBadBody(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		t.Error("Cannot create temporary file", err)
	}
	defer os.Remove(tmpFile.Name())

	// yaml test
	input := []byte(`ABC`)

	if _, err = tmpFile.Write(input); err != nil {
		t.Error("Failed to write to temporary file", err)
	}

	_, err = LoadServiceConfig(tmpFile.Name())
	if err != nil {
		t.Log("Wrong file")
	}

}

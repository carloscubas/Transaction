package config

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestLoadServiceConfig ...
func TestLoadServiceConfig(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		t.Error("Cannot create temporary file", err)
	}
	defer os.Remove(tmpFile.Name())

	// yaml test
	input := []byte(`service_name: "iam"

log_level: "DEBUG"

log_dump: false

profiler_enabled: false

swagger_enabled: true

server:
    address: ":8080"
    write_timeout: "15s"
    read_timeout: "15s"
    idle_timeout: "1m"
    shutdown_timeout: "30s"`)

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
func TestLoadServiceConfigBadPath(t *testing.T) {

	_, err := LoadServiceConfig("ABC")
	if err != nil {
		t.Log("Wrong file")
	}

}

// TestLoadServiceConfigBadBody ...
func TestLoadServiceConfigBadBody(t *testing.T) {
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

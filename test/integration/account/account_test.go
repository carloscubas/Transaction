package account

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())

	defer ts.Close()

	fmt.Printf("%s/accounts/1", ts.URL)

	resp, err := http.Get(fmt.Sprintf("%s/v1/accounts/1", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}

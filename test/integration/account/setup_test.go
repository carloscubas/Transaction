package account

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"transaction/internal/account"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func newTestServer(config account.Config, repository account.Repository) (*httptest.Server, error) {

	var (
		r              = mux.NewRouter()
		g              = gin.New()
		// s              = account.NewService(config, repository)
		accountHandler = account.NewHandler(nil, config.Logger)
	)

	account.SetRoutes(accountHandler, config, g)
	return httptest.NewServer(r), nil
}

/*
func newMockServer() (*httptest.Server, error) {
	mw := integration.NewTestMiddleware()
	l := zap.Logger{}
	r := mux.NewRouter()
	mockHandler := NewHandler(&l)
	SetMockRoutes(mockHandler, r, &mw.Middleware)
	return httptest.NewServer(r), nil
}

 */


func requestAccounts(ts *httptest.Server) (int, []byte, error) {
	accountURL := ts.URL + "/v1/accounts/1"
	rr, err := createGetRequest(accountURL)
	if err != nil {
		return -1, nil, err
	}
	defer rr.Body.Close()
	bodyBytes, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		return -1, nil, err
	}
	return rr.StatusCode, bodyBytes, nil
}

func createGetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}


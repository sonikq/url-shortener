package user

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/http/httptest"
)

func Example_handlerPingDB() {

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	srv := httptest.NewServer(mux)

	req := resty.New().R()
	req.Method = http.MethodGet

	req.URL = srv.URL + "/ping"
	req.Header.Add("Content-Type", "text/plain")
	_, err := req.Send()

	fmt.Printf("%v", err)

	// Output:
	// <nil>
}

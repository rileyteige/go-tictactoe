package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	methodGet    = "GET"
	methodPost   = "POST"
	methodPut    = "PUT"
	methodDelete = "DELETE"
)

func writeJson(w http.ResponseWriter, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func readJson(r *http.Request, v interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func errorStandardText(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func errorMethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func initRequest(f http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received %v request for %v\n", r.Method, r.URL)
		f.ServeHTTP(w, r)
	}
}

func Listen(port uint) error {
	router := buildRouter()
	s := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: initRequest(router),
	}

	return s.ListenAndServe()
}

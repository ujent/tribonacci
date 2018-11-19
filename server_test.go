package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestTribonacciRoute(t *testing.T) {

	logger := log.New(os.Stdout, "test: ", log.LstdFlags)
	s := newServer(logger)

	r, err := http.NewRequest("GET", "/tribonacci/10", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	s.http.Handler.ServeHTTP(w, r)

	resp := w.Result()

	// assert status
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Invalid status. Expected: %d, got: %d", http.StatusOK, resp.StatusCode)
	}

	// assert result
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	res := &tribonacciRS{}
	err = json.Unmarshal(body, res)
	if err != nil {
		t.Error(err)
	}

	expected := big.NewInt(44)
	if res.Result.Cmp(expected) != 0 {
		t.Errorf("Invalid result. Expected: %v, got: %v", expected, res.Result)
	}
}

var tests = []string{"/tribonacci/-10", "/tribonacci/1.2", "/tribonacci/sss"}

func TestTribonacciRoute_fail(t *testing.T) {
	logger := log.New(os.Stdout, "test: ", log.LstdFlags)
	s := newServer(logger)

	for _, route := range tests {
		r, err := http.NewRequest("GET", route, nil)
		if err != nil {
			t.Error(err)
		}
		w := httptest.NewRecorder()

		s.http.Handler.ServeHTTP(w, r)

		resp := w.Result()

		// assert status
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Invalid status. Expected: %d, got: %d", http.StatusBadRequest, resp.StatusCode)
		}
	}

}

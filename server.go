package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"tribonacci/sequence"
)

const tribonacciRoute = "/tribonacci/"

type server struct {
	http   *http.Server
	logger *log.Logger
}

type tribonacciRS struct {
	Result  *big.Int `json:"result,omitempty"`
	Message string   `json:"message,omitempty"`
}

func newServer(logger *log.Logger) *server {
	s := server{logger: logger}
	router := http.NewServeMux()
	router.HandleFunc(tribonacciRoute, s.handleTribonacci)

	s.http = &http.Server{Handler: s.loggingMiddleware(router.ServeHTTP)}
	return &s
}

func (s *server) loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next(w, r)
	}

}

func (s *server) writeJSON(w http.ResponseWriter, statusCode int, payload interface{}) {

	json, err := json.Marshal(payload)
	if err != nil {
		s.logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}

func (s *server) handleTribonacci(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	nparam := strings.TrimPrefix(r.URL.Path, tribonacciRoute)
	nparam = strings.TrimSuffix(nparam, "/")

	if len(nparam) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}

	n := new(big.Int)
	n, ok := n.SetString(nparam, 10)

	if !ok {
		s.writeJSON(w, http.StatusBadRequest,
			&tribonacciRS{
				Message: fmt.Sprintf("Wrong parameter: %s", nparam),
			})
		return
	}

	res, err := sequence.Tribonacci(n)

	if err != nil {
		if err == sequence.ErrArgNotPositive {
			s.writeJSON(w, http.StatusBadRequest,
				&tribonacciRS{
					Message: fmt.Sprintf("Wrong parameter: %s", nparam),
				})
			return
		}
	}

	s.writeJSON(w, http.StatusOK, &tribonacciRS{
		Result: res,
	})
}

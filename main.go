package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {

	logger := log.New(os.Stdout, "tribonacci: ", log.LstdFlags)

	s := newServer(logger)

	addr, ok := os.LookupEnv("TRIBONACCI_ADDR")
	if !ok {
		addr = ":8000"
	}

	var graceTimeout int
	graceTimeoutEnv, ok := os.LookupEnv("TRIBONACCI_GRACE_TIMEOUT")

	if ok {
		var err error
		graceTimeout, err = strconv.Atoi(graceTimeoutEnv)

		if err != nil {
			log.Fatalf("Invalid TRIBONACCI_GRACE_TIMEOUT value: %s", graceTimeoutEnv)
		}
	} else {
		graceTimeout = 10
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatal(err)
	}

	// gracefull shutdown
	quit := make(chan bool)
	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-shutdown

		logger.Println("Shutdown server..")

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(graceTimeout)*time.Second)
		defer cancel()

		if err := s.http.Shutdown(ctx); err != nil {
			logger.Fatalf("Unable to gracefully shutdown server. Error: %v", err)
		}

		quit <- true

	}()

	logger.Printf("Server listening on address %s", addr)

	if err := s.http.Serve(l); err != nil && err != http.ErrServerClosed {
		logger.Fatal(err)
	}

	<-quit
}

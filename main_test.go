package main

import (
	"testing"

	"bwa.com/hello/test"
)

func TestMain(m *testing.M) {
	// Queries
	queries := test.MockedQueries{}

	// App
	_ = NewApp(&queries)

	// Server
	//srv := &http.Server{
	//	Handler: a.Router,
	//	Addr:    ":3002",
	//}

	// TODO
}

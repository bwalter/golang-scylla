package main

import (
	"fmt"
	"net/http"
	"os"

	"bwa.com/hello/db"
	"github.com/jessevdk/go-flags"
)

func main() {
	// Command line args
	var opts struct {
		Addr string `short:"a" long:"addr" description:"hostname or address of the ScyllaDB node (e.g. 172.17.0.2)" required:"yes"`
		Port int    `short:"p" long:"port" description:"port of the ScyllaDB node (default: 9042)" default:"9042"`
	}
	_, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	// DB Queries
	queries, err := db.StartDbSessionAndCreateQueries(fmt.Sprintf("%s:%d", opts.Addr, opts.Port))
	if err != nil {
		panic(err)
	}

	// App
	a := NewApp(queries)

	// Server
	srv := &http.Server{
		Handler: a.Router,
		Addr:    ":3001",
	}
	srv.ListenAndServe()
}

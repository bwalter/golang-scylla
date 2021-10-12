package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"bwa.com/hello/app"
	"bwa.com/hello/db/scylla"
	"github.com/jessevdk/go-flags"
)

// @title Hello API
// @desc This is a sample project.
// Have fun using the API!
// @terms https://www.example.com/terms
//
// @contact.name Benoit Walter
// @contact.url https://github.com/bwalter
// @contact.email b.walter@meerun.de
//
// @lic.name Apache 2.0
// @lic.url https://www.apache.org/licenses/LICENSE-2.0.html
//
// @ver 1.0
// @server https://api.example.com/v1 Production API
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Command line args
	var opts struct {
		Addr string `short:"a" long:"addr" description:"hostname or address of the ScyllaDB node (e.g. 172.17.0.2)" default:"localhost"`
		Port int    `short:"p" long:"port" description:"port of the ScyllaDB node (default: 9042)" default:"9042"`
	}
	_, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	// Host and keyspace
	host := fmt.Sprintf("%s:%d", opts.Addr, opts.Port)
	keyspace := "hello"

	// Scylla keyspace
	err = scylla.CreateKeyspace(host, keyspace, false)
	if err != nil {
		panic(err)
	}

	// Scylla database
	db, err := scylla.StartSessionAndCreateDatabase(host, keyspace)
	if err != nil {
		panic(err)
	}

	// App
	a := app.NewApp(db)

	// Server
	srv := &http.Server{
		Handler: a.Router,
		Addr:    ":3001",
	}

	fmt.Printf("Listening to %s:%d...\n", "localhost", 3001)
	if err = srv.ListenAndServe(); err != nil {
		panic(err)
	}

	// Close session
	a.CloseSession()
}

module bwa.com/hello

go 1.17

require (
	github.com/gocql/gocql v1.5.0
	github.com/gorilla/mux v1.8.0 // direct
	github.com/scylladb/gocqlx v1.5.0 // direct
)

require (
	github.com/golang/snappy v0.0.1 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/jessevdk/go-flags v1.5.0 // indirect
	github.com/scylladb/go-reflectx v1.0.1 // indirect
	golang.org/x/sys v0.0.0-20210320140829-1e4c9ba3b0c4 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.5.0

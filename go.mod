module bwa.com/hello

go 1.17

require (
	github.com/gocql/gocql v1.5.0
	github.com/golang/mock v1.6.0
	github.com/gorilla/mux v1.8.0 // direct
	github.com/jessevdk/go-flags v1.5.0
	github.com/scylladb/gocqlx/v2 v2.4.0
	github.com/stretchr/testify v1.7.0
	gopkg.in/go-playground/validator.v9 v9.31.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/scylladb/go-reflectx v1.0.1 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.5.0

# golang-scylla

Useless, (almost) production-ready demo web application.

Features:
- Rest API to create, find and delete vehicles
- Dummy (mock) database

### Software Design

#### Architecture:

![architecture image](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.github.com/bwalter/golang-scylla/without-db/doc/architecture.plantuml)

#### Libraries/tools:
- HTTP server (based on net/http), see main.go
- Basic routing (based on [gorilla/mux](https://github.com/gorilla/mux)), see app/app.go
- (Basic) OpenAPI generation based on [apidoc](https://github.com/spaceavocado/apidoc), see Makefile ('apidoc' target), comments in main.go and app/app.go
- JSON validation based on [validator](https://github.com/go-playground/validator)
- Command line arguments parsing based on [go-flags](https://github.com/jessevdk/go-flags)
- Unit tests based on [testify](https://github.com/stretchr/testify)/require and [mockgen](https://github.com/golang/mock), see app/app_test.go and Makefile ('unit-tests' target)
- Integration test using httptest, see integration_tests/app_integration_test.go and Makefile ('integration-tets' target)
- Linting based on [golangci-lint](https://golangci-lint.run/), see Makefile ('check' target) and .golangci.yaml

#### Continuous integration:
- Github actions, see .github/workflows/go.yml
- Docker image generation, see Dockerfile

#### TODO:
- Database migrations
- Authentication/session
- i18n

#### Class diagram:
![classdiagram image](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.github.com/bwalter/golang-scylla/without-db/doc/classdiagram.plantuml)

### Build and run demo app

Manually:
```
$ make check
$ make
$ ./bin/hello
```

Via docker:
```
$ docker build -t hello-app .
$ docker run -t -i -p 3001:3001 -it hello-app
```

Via docker-compose:
```
$ docker-compose build
$ docker-compose up
```

### Tests

All tests:
```
$ make test
```

Unit tests only:
```
$ make unit-tests
```

Integration tests:
```
$ make integration-tests
```

If it still crashes even after testing:
try the [Rust version](https://github.com/bwalter/rust-axum-scylla) :)

### Test Rest API

Create vehicles:
```
$ curl -v -H "Accept: application/json" -H "Content-type: application/json" localhost:3001/vehicle -d '{"vin":"vin1","engine":"Combustion"}'
$ curl -v -H "Accept: application/json" -H "Content-type: application/json" localhost:3001/vehicle -d '{"vin":"vin2","engine":"Ev", ev_data: {"battery_capacity_in_kwh": 62, "soc_in_percent": 74}}
$ curl -v -H "Accept: application/json" -H "Content-type: application/json" localhost:3001/vehicle -d '{"vin":"vin3","engine":"Phev"}}'
```

Find vehicle by vin:
```
$ curl -v -H "Accept: application/json" -H "Content-type: application/json" localhost:3001/vehicle -G --data-urlencode 'vin=vin2'
```

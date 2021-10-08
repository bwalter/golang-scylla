package app

import (
	"bwa.com/hello/db"
	"bwa.com/hello/routing"

	"github.com/gorilla/mux"
)

type App struct {
	Router  *mux.Router
	queries db.Queries
}

// Create app with a router for handling requests
func NewApp(q db.Queries) App {
	r := mux.NewRouter()

	// Handle routing
	routing.Route(r, q)

	a := App{
		Router:  r,
		queries: q,
	}

	return a
}

func (a *App) CloseSession() {
	a.queries.CloseSession()
}

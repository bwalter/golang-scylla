package app

import (
	"bwa.com/hello/db"
	"bwa.com/hello/routing"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database db.Database
}

// Create app with a router for handling requests
func NewApp(db db.Database) App {
	r := mux.NewRouter()

	// Handle routing
	routing.Route(r, db)

	a := App{
		Router:   r,
		Database: db,
	}

	return a
}

func (a *App) CloseSession() {
	a.Database.CloseSession()
}

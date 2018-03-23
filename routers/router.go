package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	// Routes for the User entity
	router = SetUserRoutes(router)
	// Routes for the Task entity
	router = SetTaskRoutes()
	// Routes for the TaskNote entity
	router = SetNoteRoutes()

	return router
}

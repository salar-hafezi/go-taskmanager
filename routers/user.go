package routers

import (
	"github.com/gorilla/mux"
	"github.com/salar-hafezi/go-taskmanager/controllers"
)

const (
	USERS_REGISTER string = "/users/register"
	USERS_LOGIN    string = "/users/login"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc(USERS_REGISTER, controllers.Register).Methods("POST")
	router.HandleFunc(USERS_LOGIN, controllers.Login).Methods("POST")
	return router
}

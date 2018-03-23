package routers

import (
	"github.com/gorilla/mux"
	"github.com/salar-hafezi/go-taskmanager/common"
	"github.com/salar-hafezi/go-taskmanager/controllers"
)

const (
	CreateTask     string = "/tasks"
	UpdateTask     string = "/tasks/{id}"
	GetTasks       string = "/tasks"
	GetTaskById    string = "/tasks/{id}"
	GetTasksByUser string = "/tasks/users/{id}"
	DeleteTask     string = "/tasks/{id}"
)

func SetTaskRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc(CreateTask, common.Authorize(controllers.CreateTask)).Methods("POST")
	router.HandleFunc(UpdateTask, common.Authorize(controllers.UpdateTask)).Methods("PUT")
	router.HandleFunc(GetTasks, common.Authorize(controllers.GetTasks)).Methods("GET")
	router.HandleFunc(GetTaskById, common.Authorize(controllers.GetTaskById)).Methods("GET")
	router.HandleFunc(GetTasksByUser, common.Authorize(controllers.GetTasksByUser)).Methods("GET")
	router.HandleFunc(DeleteTask, common.Authorize(controllers.DeleteTask)).Methods("DELETE")

	return router
}

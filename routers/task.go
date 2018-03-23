package routers

import (
	"net/http"

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
	router.Handle(CreateTask, common.Authorize(http.HandlerFunc(controllers.CreateTask))).Methods("POST")
	router.Handle(UpdateTask, common.Authorize(http.HandlerFunc(controllers.UpdateTask))).Methods("PUT")
	router.Handle(GetTasks, common.Authorize(http.HandlerFunc(controllers.GetTasks))).Methods("GET")
	router.Handle(GetTaskById, common.Authorize(http.HandlerFunc(controllers.GetTaskByID))).Methods("GET")
	router.Handle(GetTasksByUser, common.Authorize(http.HandlerFunc(controllers.GetTasksByUser))).Methods("GET")
	router.Handle(DeleteTask, common.Authorize(http.HandlerFunc(controllers.DeleteTask))).Methods("DELETE")

	return router
}

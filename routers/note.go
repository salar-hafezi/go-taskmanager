package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/salar-hafezi/go-taskmanager/common"
	"github.com/salar-hafezi/go-taskmanager/controllers"
)

const (
	CreateNote     string = "/notes"
	UpdateNote     string = "/notes/{id}"
	GetNoteById    string = "/notes/{id}"
	GetNotes       string = "/notes"
	GetNotesByTask string = "/notes/tasks/{id}"
	DeleteNote     string = "/notes/{id}"
)

func SetNoteRoutes(router *mux.Router) *mux.Router {
	router.Handle(CreateNote, common.Authorize(http.HandlerFunc(controllers.CreateNote))).Methods("POST")
	router.Handle(UpdateNote, common.Authorize(http.HandlerFunc(controllers.UpdateNote))).Methods("PUT")
	router.Handle(GetNoteById, common.Authorize(http.HandlerFunc(controllers.GetNoteByID))).Methods("GET")
	router.Handle(GetNotes, common.Authorize(http.HandlerFunc(controllers.GetNotes))).Methods("GET")
	router.Handle(GetNotesByTask, common.Authorize(http.HandlerFunc(controllers.GetNotesByTask))).Methods("GET")
	router.Handle(DeleteNote, common.Authorize(http.HandlerFunc(controllers.DeleteNote))).Methods("DELETE")

	return router
}

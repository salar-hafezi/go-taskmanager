package routers

import (
	"github.com/gorilla/mux"
	"github.com/salar-hafezi/go-taskmanager/common"
	"github.com/salar-hafezi/go-taskmanager/controllers"
)

const (
	CreateNote string = "/notes"
	UpdateNote string "/notes/{id}"
	GetNoteById string = "/notes/{id}"
	GetNotes string = "/notes"
	GetNotesByTask string = "/notes/tasks/{id}"
	DeleteNote string = "/notes/{id}"
)

func SetNoteRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc(CreateNote, common.Authorize(controllers.CreateNote)).Methods("POST")
	router.HandleFunc(UpdateNote, common.Authorize(controllers.UpdateNote)).Methods("PUT")
	router.HandleFunc(GetNoteById, common.Authorize(controllers.GetNoteById)).Methods("GET")
	router.HandleFunc(GetNotes, common.Authorize(controllers.GetNotes)).Methods("GET")
	router.HandleFunc(GetNotesByTask, common.Authorize(controllers.GetNotesByTask)).Methods("GET")
	router.HandleFunc(DeleteNote, common.Authorize(controllers.DeleteNote)).Methods("DELETE")

	return router
}
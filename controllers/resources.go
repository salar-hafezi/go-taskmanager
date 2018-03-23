package controllers

import (
	"github.com/salar-hafezi/go-taskmanager/models"
)

type (
	// POST /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}

	// POST /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}

	// response for authorized user POST /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}

	// POST/PUT /tasks
	// GET /tasks/id
	TaskResource struct {
		Data models.Task `json:"data"`
	}

	// GET /tasks
	TasksResource struct {
		Data []models.Task `json:"data"`
	}

	// POST/PUT /notes
	NoteResource struct {
		Data NoteModel `json:"data"`
	}

	// GET /notes
	// /notes/tasks/id
	NotesResource struct {
		Data []models.TaskNote `json:"data"`
	}

	// authentication model
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// authorized user with access token model
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}

	// TaskNote model
	NoteModel struct {
		TaskId      string `json:"taskid"`
		Description string `json:"description"`
	}
)

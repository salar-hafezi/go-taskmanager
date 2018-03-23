package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResourse struct {
		Data appError `json:"data"`
	}
	configuration struct {
		Server, MongoDBHost, DBUser, DBPwd, Database string
		LogLevel                                     int
	}
)

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResourse{Data: errObj}); err == nil {
		w.Write(j)
	}
}

// AppConfig holds the config values from config.json
var AppConfig configuration

// Initialize AppConfig
func initConfig() {
	loadAppConfig()
}

// Read config.json and decode to AppConfig
func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}

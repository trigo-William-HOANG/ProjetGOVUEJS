package api

import (
	"encoding/json"
	"net/http"
)

type AppDataParams struct {
	Id int
}

type AppDataResponse struct {
	Code        int
	Title       string
	Description string
	Link        string
	Image       string
}

type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, code int, message string) {
	resp := Error{Code: code, Message: message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, http.StatusBadRequest, err.Error())
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, http.StatusInternalServerError, "Internal Server Error")
	}
)

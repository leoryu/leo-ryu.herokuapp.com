package render

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func ErrorCode(w http.ResponseWriter, err error, status int) {
	JSON(w, Error{Message: err.Error()}, status)
}

func InternalError(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 500)
}

func NotImplemented(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 501)
}

func NotFound(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 404)
}

func Unauthorized(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 401)
}

func BadRequest(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 400)
}

func JSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.Encode(v)
}

package views

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/BRO3886/meetings-api/pkg"
)

//ErrView for wrapping error
type ErrView struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//noinspection ALL
var (
	ErrMethodNotAllowed = errors.New("Error: Method is not allowed")
	ErrInvalidToken     = errors.New("Error: Invalid Authorization token")
	ErrUserExists       = errors.New("User already exists")
)

//ErrHTTPStatusMap maps pkg error to view error with http status code
var ErrHTTPStatusMap = map[string]int{
	pkg.ErrNotFound.Error():     http.StatusNotFound,
	pkg.ErrInvalidSlug.Error():  http.StatusBadRequest,
	pkg.ErrExists.Error():       http.StatusConflict,
	pkg.ErrNoContent.Error():    http.StatusNotFound,
	pkg.ErrDatabase.Error():     http.StatusInternalServerError,
	pkg.ErrUnauthorized.Error(): http.StatusUnauthorized,
	pkg.ErrForbidden.Error():    http.StatusForbidden,
	ErrMethodNotAllowed.Error(): http.StatusMethodNotAllowed,
	ErrInvalidToken.Error():     http.StatusBadRequest,
	ErrUserExists.Error():       http.StatusConflict,
}

//Wrap func to send error
func Wrap(err error, w http.ResponseWriter) {
	msg := err.Error()
	code := ErrHTTPStatusMap[msg]

	// If error code is not found
	// like a default case
	if code == 0 {
		code = http.StatusInternalServerError
	}

	w.WriteHeader(code)

	errView := ErrView{
		Message: msg,
		Status:  code,
	}
	log.Printf("Error: " + msg + " Code: " + fmt.Sprint(code))

	json.NewEncoder(w).Encode(errView)
}

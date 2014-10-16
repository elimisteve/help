// Steve Phillips / elimisteve
// 2014.10.15

package help

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
)

func WriteError(w http.ResponseWriter, msg string, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	writeJSONError(w, msg, statusCode)
	log.Printf("WriteError: %d: %s\n", statusCode, msg)

	if statusCode == 500 {
		trace := make([]byte, 1024)
		runtime.Stack(trace, true)
		log.Printf("Stacktrace:\n %s", trace)
	}
}

func writeJSONError(w http.ResponseWriter, msg string, statusCode int) {
	// TODO(elimisteve): Dynamically set the status field later
	e := jsonError{Error: msg, StatusCode: statusCode, Status: "error"}
	errJSON, _ := json.Marshal(&e)
	http.Error(w, string(errJSON), statusCode)
}

type jsonError struct {
	Error      string `json:"error"`
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
}

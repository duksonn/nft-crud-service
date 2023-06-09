package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type errorResponse struct {
	Code     string
	Messages []string
}

/** renders obj with 200 status code */
func OK(w http.ResponseWriter, r *http.Request, obj interface{}) {
	Render(w, r, obj, http.StatusOK)
}

/** response as json */
func BadRequest(w http.ResponseWriter, r *http.Request, code string, messages ...string) {
	err := &errorResponse{
		Code:     code,
		Messages: messages,
	}
	Render(w, r, err, http.StatusBadRequest)
}

/** response as json */
func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	Render(w, r, &errorResponse{
		Code:     "INTERNAL_SERVER_ERROR",
		Messages: []string{err.Error()},
	}, http.StatusInternalServerError)
}

/** response as json */
func Render(w http.ResponseWriter, r *http.Request, obj interface{}, status int) {
	js, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

/** parses a int parameter from url */
func GetIntParam(r *http.Request, key string, defaultValue int) (int, error) {
	keys, ok := r.URL.Query()[key]

	if !ok || len(keys[0]) < 1 {
		return defaultValue, nil
	}

	value, err := strconv.ParseInt(keys[0], 10, 0)
	if err != nil {
		return defaultValue,
			fmt.Errorf("server: %s is not a valid int value for %s", keys[0], key)
	}

	return int(value), nil
}

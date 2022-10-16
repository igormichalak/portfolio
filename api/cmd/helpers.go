package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func loadEnvInt(dst *int, key string, defaultVal int, override bool) {
	if *dst != 0 && !override {
		return
	}
	val, found := os.LookupEnv(key)
	if !found {
		if *dst == 0 {
			*dst = defaultVal
		}
		return
	}
	parsedVal, err := strconv.Atoi(val)
	if err != nil {
		if *dst == 0 {
			*dst = defaultVal
		}
		return
	}
	*dst = parsedVal
}

func loadEnvString(dst *string, key string, defaultVal string, override bool) {
	if *dst != "" && !override {
		return
	}
	val, found := os.LookupEnv(key)
	if !found {
		if *dst == "" {
			*dst = defaultVal
		}
		return
	}
	*dst = val
}

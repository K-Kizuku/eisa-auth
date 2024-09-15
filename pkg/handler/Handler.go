package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/K-Kizuku/eisa-auth/pkg/errors"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		appErr, ok := err.(*errors.Error)
		if ok {
			w.WriteHeader(appErr.Status)
			http.Error(w, appErr.Error(), appErr.Status)
			fmt.Println(appErr.StackTrace)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			slog.Error("Recovered from a panic", "unknown errror", err)
		}
	}
}

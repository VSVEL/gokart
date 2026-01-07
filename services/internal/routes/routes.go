package routes

import (
	"net/http"

	"gokart/internal/handler"
)

func SetupRoutes(evtHandler *handler.EventHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/events", evtHandler)
	return mux
}

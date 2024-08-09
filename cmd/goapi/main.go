package main

import (
	"fmt"
	"net/http"

	"github.com/Oriseer/testruns/internal/handler"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handler.Handler(r)
	fmt.Println("Starting go api....")

	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		log.Error(err)
		return
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tonus-gps-tracker/server/utils"
)

type HttpServer struct {
	Port    int
	Tracker *Tracker
}

func (server *HttpServer) setup() {
	server.Port = utils.StringToInt(utils.GetEnv("HTTP_SERVER_PORT"))
	server.Tracker = NewTracker()
}

func (server *HttpServer) Run() {
	log.Println("[INFO][API] Server started")

	server.setup()

	mux := mux.NewRouter()
	mux.HandleFunc("/save", server.Tracker.Save).Methods(http.MethodPost)
	mux.HandleFunc("/health", server.Tracker.Health).Methods(http.MethodGet)

	err := http.ListenAndServe(
		fmt.Sprintf(":%d", server.Port),
		mux,
	)

	if err != nil {
		log.Fatalf("[ERROR] HttpServer_Run, http.ListenAndServe: %s\n", err)
	}
}

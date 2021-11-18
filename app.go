package main

import "github.com/gorilla/mux"

type App struct {
	router *mux.Router
	port   string
}

func NewApp(port string) *App {
	r := mux.NewRouter()
	return &App{
		port:   port,
		router: r,
	}
}

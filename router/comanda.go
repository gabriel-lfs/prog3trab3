package router

import (
	"github.com/gabriel-lfs/prog3trab3/handlers"
	"github.com/gorilla/mux"
)

func comandaRouting(router mux.Router) mux.Router {
	router.HandleFunc("/comanda/{id}", handlers.GetComandaHandler).Methods("GET")
	router.HandleFunc("/comanda", handlers.GetComandasHandler).Methods("GET")
	router.HandleFunc("/comanda", handlers.CreateComandaHandler).Methods("POST")
	router.HandleFunc("/comanda{id}", handlers.DeleteComandaHandler).Methods("DELETE")
	router.HandleFunc("/comanda/{id}", handlers.UpdateComandaHandler).Methods("PUT")
	return router
}

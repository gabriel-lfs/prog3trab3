package router

import (
	"github.com/gorilla/mux"
	"go-prog/handlers"
)

func comandaRouting(router mux.Router) mux.Router {
	router.HandleFunc("/comanda/{id}", handlers.GetComandaHandler).Methods("GET")
	router.HandleFunc("/comanda", handlers.GetComandasHandler).Methods("GET")
	router.HandleFunc("/comanda", handlers.CreateComandaHandler).Methods("POST")
	router.HandleFunc("/comanda{id}", handlers.DeleteComandaHandler).Methods("DELETE")
	router.HandleFunc("/comanda/{id}", handlers.UpdateComandaHandler).Methods("PUT")
	return router
}

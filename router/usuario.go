package router

import (
	"github.com/gorilla/mux"
	"go-prog/handlers"
)

func usuarioRouting(router mux.Router) mux.Router {
	router.HandleFunc("/usuarios/{id}", handlers.GetUsuarioHandler).Methods("GET")
	router.HandleFunc("/usuarios", handlers.GetUsuariosHandler).Methods("GET")
	router.HandleFunc("/usuarios", handlers.CreateUsuarioHandler).Methods("POST")
	router.HandleFunc("/usuarios{id}", handlers.DeleteUsuarioHandler).Methods("DELETE")
	router.HandleFunc("/usuarios/{id}", handlers.UpdateUsuarioHandler).Methods("PUT")

	return router
}

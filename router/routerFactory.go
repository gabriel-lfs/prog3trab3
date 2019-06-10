package router

import (
    "github.com/gorilla/mux"
)

func StartRouting() *mux.Router {
    router := mux.NewRouter()
    router.Headers("Content-Type", "application/json", "Access-Control-Allow-Origin", "*")
    *router = comandaRouting(*router)
    *router = usuarioRouting(*router)

    return router
}
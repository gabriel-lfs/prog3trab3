package main

import (
    "github.com/gabriel-lfs/prog3trab3/database"
    "github.com/gabriel-lfs/prog3trab3/router"
    "log"
    "net/http"
)

func main() {
    database.Migrate()
    request := router.StartRouting()
    log.Fatal(http.ListenAndServe(":8080", request))
}
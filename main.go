package main

import (
    "go-prog/database"
    "go-prog/router"
    "log"
    "net/http"
)

func main() {
    database.Migrate()
    request := router.StartRouting()
    log.Fatal(http.ListenAndServe(":8080", request))
}
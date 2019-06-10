package handlers

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "go-prog/database"
    "net/http"
    "strconv"
)

func CreateComandaHandler(w http.ResponseWriter, r *http.Request) {
    var comanda database.Comanda
    err := json.NewDecoder(r.Body).Decode(&comanda)

    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    database.CreateComanda(comanda)
}

func UpdateComandaHandler(w http.ResponseWriter, r *http.Request) {
    urlParams := mux.Vars(r)
    idStr := urlParams["id"]
    id, err := strconv.Atoi(idStr)

    if err != nil {
        http.Error(w, err.Error(), 500)
    }

    var comanda database.Comanda
    err = json.NewDecoder(r.Body).Decode(&comanda)

    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    comanda.ID = uint(id)

    database.UpdateComanda(comanda, uint(id))
}

func DeleteComandaHandler(w http.ResponseWriter, r *http.Request) {
    urlParams := mux.Vars(r)
    idStr := urlParams["id"]
    id, err := strconv.Atoi(idStr)

    if err != nil {
        http.Error(w, err.Error(), 500)
    }

    database.DeleteComanda(uint(id))
}

func GetComandaHandler(w http.ResponseWriter, r *http.Request) {
    urlParams := mux.Vars(r)
    idStr := urlParams["id"]
    id, err := strconv.Atoi(idStr)

    if err != nil {
        http.Error(w, err.Error(), 500)
    }

    comanda := database.GetComanda(id)

    jsonComanda, _ := json.Marshal(comanda)

    _ , err = w.Write(jsonComanda)

    if err != nil {
        http.Error(w, err.Error(), 500)
    }
}

func GetComandasHandler(w http.ResponseWriter, r *http.Request) {
    comandas := database.GetComandas()

    jsonComanda, err := json.Marshal(&comandas)

    if err != nil {
        http.Error(w, err.Error(), 500)
    }
    _, _ = w.Write(jsonComanda)
}

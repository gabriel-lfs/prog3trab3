package handlers

import (
    "encoding/json"
    "github.com/gabriel-lfs/prog3trab3/database"
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
)

func CreateUsuarioHandler(w http.ResponseWriter, r *http.Request) {
    var usuario database.Usuario
    err := json.NewDecoder(r.Body).Decode(&usuario)

    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    database.CreateUsuario(usuario)
}

func UpdateUsuarioHandler(w http.ResponseWriter, r *http.Request) {
    urlParams := mux.Vars(r)
    idStr := urlParams["id"]
    id, err := strconv.Atoi(idStr)

    if err != nil {
        http.Error(w, err.Error(), 500)
    }

    var usuario database.Usuario
    err = json.NewDecoder(r.Body).Decode(&usuario)

    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    usuario.ID = uint(id)

    database.UpdateUsuario(usuario, uint(id))
}

func DeleteUsuarioHandler(w http.ResponseWriter, r *http.Request) {
    urlParams := mux.Vars(r)
    idStr := urlParams["id"]
    id, err := strconv.Atoi(idStr)

    if err != nil {
        http.Error(w, err.Error(), 500)
    }

    database.DeleteUsuario(uint(id))
}

func GetUsuarioHandler(w http.ResponseWriter, r *http.Request) {
    urlParams := mux.Vars(r)
    idStr := urlParams["id"]
    id, err := strconv.Atoi(idStr)

    if err != nil {
        http.Error(w, err.Error(), 500)
    }

    usuario := database.GetUsuario(id)

    jsonUsuario, _ := json.Marshal(usuario)

    _ , err = w.Write(jsonUsuario)

    if err != nil {
        http.Error(w, err.Error(), 500)
    }
}

func GetUsuariosHandler(w http.ResponseWriter, r *http.Request) {
    usuarios := database.GetUsuarios()

    jsonUsuario, err := json.Marshal(&usuarios)

    if err != nil {
        http.Error(w, err.Error(), 500)
    }
    _, _ = w.Write(jsonUsuario)
}
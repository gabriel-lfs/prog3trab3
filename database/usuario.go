package database

import (
    "github.com/jinzhu/gorm"
)

type Usuario struct {
    gorm.Model
    Email string `json:"email"`
    Senha string `json:"senha"`
}


func CreateUsuario(usuario Usuario) {
    db := openDBConnection()
    defer db.Close()

    db.Create(&usuario)
}

func UpdateUsuario(usuario Usuario, id uint) Usuario {
    db := openDBConnection()
    defer db.Close()

    var newUsuario Usuario
    newUsuario.ID = id

    db.Model(&newUsuario).Updates(&usuario)

    return newUsuario
}

func DeleteUsuario(id uint) {
    db := openDBConnection()
    defer db.Close()
    var usuario Usuario
    usuario.ID = id

    db.Delete(&usuario)
}

func GetUsuario(id int) Usuario {
    db := openDBConnection()
    defer db.Close()
    var usuario Usuario

    db.First(&usuario, "id=?", id)

    return usuario
}

func GetUsuarios() []Usuario {
    db := openDBConnection()
    defer db.Close()

    var usuarios []Usuario
    db.Find(&usuarios)

    return usuarios
}
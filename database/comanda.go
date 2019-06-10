package database

import (
    "github.com/jinzhu/gorm"
)

type Comanda struct {
    gorm.Model
    Usuario    Usuario
    UsuarioID  int `json:"idusuario"`
    Produtos   string `json:"produtos"`
    ValorTotal string `json:"valortotal"`
}

func CreateComanda(comanda Comanda) {
    db := openDBConnection()
    defer db.Close()

    db.Create(&comanda)
}

func UpdateComanda(comanda Comanda, id uint) Comanda {
    db := openDBConnection()
    defer db.Close()

    var newComanda Comanda
    newComanda.ID = id

    db.Model(&newComanda).Updates(&comanda)

    return newComanda
}

func DeleteComanda(id uint) {
    db := openDBConnection()
    defer db.Close()
    var comanda Comanda
    comanda.ID = id

    db.Delete(&comanda)
}

func GetComanda(id int) Comanda {
    db := openDBConnection()
    defer db.Close()
    var comanda Comanda

    db.First(&comanda, "id=?", id)

    return comanda
}

func GetComandas() []Comanda {
    db := openDBConnection()
    defer db.Close()

    var comandas []Comanda
    db.Find(&comandas)

    return comandas
}
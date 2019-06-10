package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "log"
    "os"
)

func openDBConnection() *gorm.DB {
    db, err := gorm.Open("postgres",
        "host=" + os.Getenv("DB_HOST") +
        " port=" + os.Getenv("DB_PORT") +
        " user=" + os.Getenv("DB_USER") +
        " dbname=" + os.Getenv("DB_NAME") +
        " password=" + os.Getenv("DB_PASSWORD") +
        " sslmode=disable")

    if err != nil {
        log.Fatal(err.Error())
    }

    return db
}

func Migrate() {
    db := openDBConnection()

    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(&Comanda{}, &Usuario{})
    db.Model(&Comanda{}).AddForeignKey("usuario_id",
        "comandas(id)",
        "CASCADE",
        "CASCADE")
}
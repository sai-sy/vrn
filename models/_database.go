package model

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)


func Database() (*gorm.DB, error) {

    db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})

    if err != nil {
        log.Fatal(err.Error())
    }

    if err = db.AutoMigrate(&Person{}); err != nil {
        log.Println(err)
    }


    return db, err

}
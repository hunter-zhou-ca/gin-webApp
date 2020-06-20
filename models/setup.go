// models/setup.go

package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB
var err error

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "test.db")

  	if err != nil {
    	panic("Failed to connect to database!")
  	}

  	database.AutoMigrate(&Person{})

  	DB = database

  	//p1 := Person{ID: 1, Name: "John", Age: 31, City: "New York"}
 	//p2 := Person{ID: 2, Name: "Doe", Age: 22, City: "Vancouver"}

 	//database.Create(&p1)
 	//database.Create(&p2)

}
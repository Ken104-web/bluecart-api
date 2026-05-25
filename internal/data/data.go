package data

import(
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Data(){
	db, err := gorm.Open(
		sqlite.Open("app.db"),
		&gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}

	DB = db
}



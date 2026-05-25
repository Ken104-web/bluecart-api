package main

import (
	"bluecart-api/internal/data"
	"bluecart-api/internal/model"
	"log"
)

func main(){

	// call that data func
	data.Data()
	err := data.DB.AutoMigrate(model.User{}, &model.SearchHistory{})
	if err != nil {
		log.Fatal(err)
	}
}




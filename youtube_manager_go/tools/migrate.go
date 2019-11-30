package main

import (
	"github.com/s-shiki/youtube_api_server/youtube_manager_go/databases"
	"github.com/s-shiki/youtube_api_server/youtube_manager_go/models"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := databases.Connect()
	defer db.Close()

	if err != nil {
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&models.User{})
	db.Debug().AutoMigrate(&models.Favorite{})
}
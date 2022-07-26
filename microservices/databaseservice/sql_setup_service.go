package databaseservice

import (
	"fmt"
	"github.com/admin-agora/backend/sql/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// TODO:
// TODO:

// TODO: Refactor this into a gin context microservice

// TODO:
// TODO:


func SQLSetUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println(dosn)
	gormDB, sqlErr := gorm.Open(mysql.Open(dosn), &gorm.Config{})
	if sqlErr != nil {
		log.Fatal("things went poorly")
	}

	// Get the SQL tables we want
	var comments entity.AgoraComment
	var agoraThreads entity.AgoraThread
	var contributor entity.Contributor
	var topics entity.Topic
	var users entity.AgoraUser
	var watcher entity.Watcher


	if commentsErr := comments.CreateTable(gormDB); commentsErr != nil {
		log.Println("could not make the comments: ", commentsErr)
	}

	if threadsErr := agoraThreads.CreateTable(gormDB); threadsErr != nil {
		log.Println("could not make the threadsErr: ", threadsErr)
	}

	if contributorErr := contributor.CreateTable(gormDB); contributorErr != nil {
		log.Println("could not make the contributorErr: ", contributorErr)
	}

	if topicsErr := topics.CreateTable(gormDB); topicsErr != nil {
		log.Println("could not make the topicsErr: ", topicsErr)
	}
	if userErr := users.CreateTable(gormDB); userErr != nil {
		log.Println("could not make the userErr: ", userErr)
	}

	if watcherErr := watcher.CreateTable(gormDB); watcherErr != nil {
		log.Println("could not make the watcherErr: ", watcherErr)
	}
}

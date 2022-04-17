package cloudfunctions

import (
	"fmt"
	"github.com/admin-agora/backend/sql/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)



func SQLSetUp(w http.ResponseWriter, r *http.Request) {
	dosn := "agora_mysql_admin:DigitalMcDonalds$3.21@tcp(agora-mysql-dev.mysql.database.azure.com)/agora?charset=utf8mb4&parseTime=true"
	fmt.Println(dosn)
	gormDB, sqlErr := gorm.Open(mysql.Open(dosn), &gorm.Config{})
	if sqlErr != nil {
		log.Fatal("things went poorly")
	}

	// Get the SQL tables we want
	var comments entity.AgoraComment
	var agoraThreads entity.AgoraThread
	var contributor entity.Contribute
	var topics entity.Topic
	var users entity.User
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

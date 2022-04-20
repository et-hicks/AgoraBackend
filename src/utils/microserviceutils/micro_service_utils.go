package microserviceutils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type HTTPError struct {
	Message string `json:"message"`
}

func GeneralInit() (*gorm.DB, *gin.Engine, error) {
	dosn := "agora_mysql_admin:DigitalMcDonalds$3.21@tcp(agora-mysql-dev.mysql.database.azure.com)/agora?charset=utf8mb4&parseTime=true"
	fmt.Println(dosn)
	gormDb, sqlErr := gorm.Open(mysql.Open(dosn), &gorm.Config{})

	// and the router
	router := gin.Default()

	if sqlErr != nil {
		return nil, nil, sqlErr
	}

	return gormDb, router, nil
}

func HardReset(*gin.Context) {
	// TODO: Learn how to gracefully shut down/restart the server
	log.Println("Ungraceful shutdown of the server")
	log.Fatal("Resetting server")
}

func BadHTTP(err error) HTTPError {
	return HTTPError{Message: err.Error()}
}


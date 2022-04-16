package posting

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type PostProcessing struct {
	db *gorm.DB
	router *gin.Engine
}

func (p *PostProcessing) Init() error {
	dosn := "agora_mysql_admin:DigitalMcDonalds$3.21@tcp(agora-mysql-dev.mysql.database.azure.com)/agora?charset=utf8mb4&parseTime=true"
	fmt.Println(dosn)
	gormDb, sqlErr := gorm.Open(mysql.Open(dosn), &gorm.Config{})

	// Assign the SQL connection to
	p.db = gormDb

	p.router = gin.Default()

	if sqlErr != nil {
		return sqlErr
	}

	return nil
}

func (p *PostProcessing) Run() error {

	p.router.GET("/hiThere/:name", hiHandler)



	routerErr := p.router.Run(":8080")
	if routerErr != nil {
		log.Fatal(routerErr)
	}
	return nil
}

func hiHandler(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "hello %s", name)
}



package commentprocessing

import (
	"fmt"
	"github.com/admin-agora/backend/messages"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type CommentPostProcessing struct {
	db     *gorm.DB
	router *gin.Engine
}

func (p *CommentPostProcessing) RunningInstance() *gorm.DB {
	return p.db
}

func (p *CommentPostProcessing) Init() error {
	dosn := "agora_mysql_admin:DigitalMcDonalds$3.21@tcp(agora-mysql-dev.mysql.database.azure.com)/agora?charset=utf8mb4&parseTime=true"
	fmt.Println(dosn)
	gormDb, sqlErr := gorm.Open(mysql.Open(dosn), &gorm.Config{})

	// Assign the SQL connection
	p.db = gormDb

	// and the router
	p.router = gin.Default()

	if sqlErr != nil {
		return sqlErr
	}

	return nil
}

func (p *CommentPostProcessing) Run() {

	// json decode to protobuf
	// persist comment
	// breakdown

	p.router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	p.router.GET("/ethan/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	p.router.GET("/name", func(c *gin.Context) {
		name := os.Getenv("K_SERVICE")
		c.String(http.StatusOK, "microservice name %s", name)
	})

	p.router.POST("/user", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		commentPosted := &messages.CommentPosted{}
		if unmarshallErr := protojson.Unmarshal(body, commentPosted); unmarshallErr != nil {
			log.Println("error in marshalling: ", unmarshallErr)
		}
		log.Println(commentPosted.Comment)
		bits, _ := protojson.Marshal(commentPosted)
		c.String(http.StatusOK, string(bits))
	})

	p.router.DELETE("/user", func (c *gin.Context) {
		log.Fatal("shutting down server")
	})
	engineErr := p.router.Run(":8080")
	if engineErr != nil {
		log.Fatal("cannot run the microservice CommentPosting")
	}
}

func Service() {
	var p CommentPostProcessing
	if initErr := p.Init(); initErr != nil {
		log.Fatal("error in init: ", initErr)
	}
	p.Run()
}




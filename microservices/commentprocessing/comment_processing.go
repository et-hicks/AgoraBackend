package commentprocessing

import (
	"github.com/admin-agora/backend/messages"
	"github.com/admin-agora/backend/src/utils/microserviceutils"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

type CommentPostProcessing struct {
	db     *gorm.DB
	router *gin.Engine
}

func (p *CommentPostProcessing) RunningInstance() *gorm.DB {
	return p.db
}

func (p *CommentPostProcessing) Init() error {

	db, r, err := microserviceutils.GeneralInit()

	if err != nil {
		return err
	}
	p.db = db
	p.router = r
	return nil
}

func (p *CommentPostProcessing) Run() {

	// Create new comment on post

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

	p.router.DELETE("/reset", func (c *gin.Context) {
		log.Fatal("Resetting server")
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




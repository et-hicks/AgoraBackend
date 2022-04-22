package commentprocessing

import (
	"github.com/admin-agora/backend/src/utils/microserviceutils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
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

	p.router.POST("/comment", CreateComment(p.db))

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




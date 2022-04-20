package threadprocessing

import (
	"github.com/admin-agora/backend/src/utils/microserviceutils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type ThreadProcessing struct {
	db     *gorm.DB
	router *gin.Engine
}

func (t *ThreadProcessing) RunningInstance() *gorm.DB {
	// Is this even required? I dont think so??
	return t.db
}

func (t *ThreadProcessing) Init() error {
	db, r, err := microserviceutils.GeneralInit()

	if err != nil {
		return err
	}
	t.db = db
	t.router = r
	return nil
}

func (t *ThreadProcessing) Run() {

	// Create new Thread
	t.router.POST("/create", CreateThread(t.db))

	// TODO: Learn a way for graceful shutdown
	t.router.DELETE("/reset", microserviceutils.HardReset)

	engineErr := t.router.Run(":8080")
	if engineErr != nil {
		log.Fatal("cannot run the microservice CommentPosting")
	}
}

func Service() {
	var t ThreadProcessing
	if initErr := t.Init(); initErr != nil {
		log.Fatal("error in init: ", initErr)
	}
	t.Run()
}


package userprocessing

import (
	"github.com/admin-agora/backend/src/utils/microserviceutils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type UserProcessing struct {
	db     *gorm.DB
	router *gin.Engine
}

func (u *UserProcessing) Init() error {
	db, r, err := microserviceutils.GeneralInit()

	if err != nil {
		return err
	}
	u.db = db
	u.router = r
	return nil
}

func (u *UserProcessing) Run() {

	u.router.POST("/login", Login(u.db))
	u.router.POST("/create", CreateUser(u.db))

	// TODO: change password

	engineErr := u.router.Run(":8080")
	if engineErr != nil {
		log.Fatal("cannot run the microservice CommentPosting")
	}

}

func Service() {
	var u UserProcessing
	if initErr := u.Init(); initErr != nil {
		log.Fatal("error in init: ", initErr)
	}
	u.Run()
}

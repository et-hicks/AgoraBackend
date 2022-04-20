package threadprocessing

import (
	"errors"
	"github.com/admin-agora/backend/messages"
	"github.com/admin-agora/backend/sql/entity"
	"github.com/admin-agora/backend/src/utils/microserviceutils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateThread(db *gorm.DB) func(ctx *gin.Context) {
	log.Println("I cannot believe this actually worked")

	return func (ctx *gin.Context) {
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		newThread := &messages.ThreadInfo{}
		if unmarshallErr := protojson.Unmarshal(body, newThread); unmarshallErr != nil {
			log.Println("error in marshalling: ", unmarshallErr)
		}
		if err := nilChecks(newThread); err != nil {
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(err))
			return
		}
		agoraThread, err := createEntity(newThread, db)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(err))
			return
		}
		log.Println("created new thread with title: ", agoraThread.Title, "and id: ", agoraThread.ID)
		ctx.JSON(http.StatusOK, agoraThread)
	}
}



func createEntity(info *messages.ThreadInfo, db *gorm.DB) (*entity.AgoraThread, error) {

	// TODO: implement access level and all that jazz

	//var user entity.User
	//db.First(&user, info.AuthorID) // DB Access
	uid, uidErr := uuid.NewUUID()
	if uidErr != nil {
		log.Fatal("could not generate thread uuid: ", uidErr)
		return nil, uidErr
	}

	thread := entity.AgoraThread{
		Title:      info.Title,
		CreatorID: uint(info.AuthorID),
		Likes:      0,
		Dislikes:   0,
		Clicks:     0,
		Watchers:   0,
		UrlUUID:    uid.String(),
		ImageURL:   info.ImageURL,
		Public:   info.PublicityLevel,
		IsReported: false,
	}
	result := db.Create(&thread) // DB Access
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println(result.RowsAffected)
	return &thread, nil
}

func nilChecks(info *messages.ThreadInfo) error {
	// need some basic things for new thread
	// title
	// public level
	// creator
	if info.Title == "" {
		return errors.New("title missing")
	}
	if unknown := messages.PublicityLevel.Enum(0); info.PublicityLevel == *unknown {
		return errors.New("public access missing")
	}

	if info.AuthorID == 0 {
		return errors.New("authorID missing")
	}

	return nil
}


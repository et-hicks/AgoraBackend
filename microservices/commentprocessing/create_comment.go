package commentprocessing

import (
	"github.com/admin-agora/backend/messages"
	"github.com/admin-agora/backend/sql/entity"
	"github.com/admin-agora/backend/src/utils/microserviceutils"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateComment(db *gorm.DB) func(ctx *gin.Context) {
	log.Println("I cannot believe this actually worked")

	return func (ctx *gin.Context) {
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		newComment := &messages.CommentPosted{}
		if unmarshallErr := protojson.Unmarshal(body, newComment); unmarshallErr != nil {
			log.Println("error in marshalling: ", unmarshallErr)
		}
		if err := nilChecks(newComment); err != nil {
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(err))
			return
		}
		agoraComment, err := createEntity(newComment, db)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(err))
			return
		}
		log.Println("created new comment with id: ", agoraComment.ID)
		ctx.JSON(http.StatusOK, agoraComment)
	}
}

func createEntity(comment *messages.CommentPosted, db *gorm.DB) (*entity.AgoraComment, error) {
	// TODO: impl
	return nil, nil
}

func nilChecks(comment *messages.CommentPosted) error {
	// TODO: impl
	return nil
}

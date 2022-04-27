package commentprocessing

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

func CreateComment(db *gorm.DB) func(ctx *gin.Context) {
	// log.Println("I cannot believe this actually worked")
	// I can do additional processing here

	// TODO: check that the user can actually comment on the thread
	// TODO: check that the thread actually exists

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

	uid, uidErr := uuid.NewUUID()
	if uidErr != nil {
		log.Fatal("could not generate thread uuid: ", uidErr)
		return nil, uidErr
	}

	if comment.ParentCommentUUID != "" {
		if parentError := confirmParentComment(comment.ParentCommentUUID, db); parentError != nil {
			return nil, parentError
		}
	}


	commentEntity := entity.AgoraComment{
		AuthorID: uint(comment.AuthorID),
		ParentCommentUUID: comment.ParentCommentUUID,
		Comment:           comment.Comment,
		IsEdited:          false,
		IsReported:        false,
		Likes:             0,
		Dislikes:          0,
		UUID: uid.String(),
		ThreadID: uint(comment.ThreadID),
	}

	result := db.Create(&commentEntity)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println(result.RowsAffected)
	return &commentEntity, nil
}

func confirmParentComment(parentUUID string, db *gorm.DB) error {
	var parentComment entity.AgoraComment
	result := db.First(&parentComment, "parent_comment_uuid = ?", parentUUID) // DB Access
	if result.Error != nil {
		return result.Error
	}
	if parentComment.ID == 0 {
		return errors.New("cannot find the original parent comment ID")
	}
	return nil
}

func nilChecks(comment *messages.CommentPosted) error {

	if comment.Comment == "" {
		return errors.New("No Comment to save in payload")
	}
	if comment.AuthorID == 0 {
		return errors.New("no author id in the comment")
	}

	if comment.ThreadID == 0 {
		return errors.New("must have  thread id")
	}

	return nil
}

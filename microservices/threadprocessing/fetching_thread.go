package threadprocessing

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

func FetchThreads(db *gorm.DB) func(ctx *gin.Context) {
	log.Println("I cannot believe this actually worked")

	return func (ctx *gin.Context) {
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		request := &messages.ThreadPageRequest{}
		if unmarshallErr := protojson.Unmarshal(body, request); unmarshallErr != nil {
			log.Println("error in marshalling: ", unmarshallErr)
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(unmarshallErr))
			return
		}
		var threads []entity.AgoraThread
		db.Limit(10).First(&threads, "created_at < FROM_UNIXTIME(?)", request.CreatedBefore)


		ctx.JSON(http.StatusOK, threads)
	}
}

func createInfoProto(threadEntity entity.AgoraThread) (*messages.ThreadDisplayInfo, error) {

	oneThread := messages.ThreadDisplayInfo{
		Title:              threadEntity.Title,
		Description:        threadEntity.Description,
		CreatedAt:          threadEntity.CreatedAt.UnixMilli(),
		ImageURL:           threadEntity.ImageURL,
		Creator:            nil,
		Contributors:       nil,
		Watchers:           0,
		Comments:           0,
		ThreadContributors: 0,
	}

	return &oneThread, nil

}

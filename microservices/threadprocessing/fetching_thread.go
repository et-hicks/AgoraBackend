package threadprocessing

import (
	"github.com/admin-agora/backend/messages"
	"github.com/admin-agora/backend/sql/entity"
	"github.com/admin-agora/backend/sql/queries"
	"github.com/admin-agora/backend/src/utils/converters"
	"github.com/admin-agora/backend/src/utils/microserviceutils"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func FetchThreadWhole(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}

func FetchThreadDisplayInfo(db *gorm.DB) func(ctx *gin.Context) {

	return func(ctx *gin.Context) {
		threadID := ctx.Param("threadID")
		u64, err := strconv.ParseUint(threadID, 10, 64)

		if err != nil {
			log.Println("error in string convert: ", err)
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(err))
			return
		}

		proto, _ := createThreadDisplayInfo(db, uint(u64))
		ctx.JSON(http.StatusOK, proto)
	}
}

// FetchThreads
// TODO: expand this with functionality
func FetchThreads(db *gorm.DB) func(ctx *gin.Context) {
	log.Println("I cannot believe this actually worked")

	return func (ctx *gin.Context) {
		body, _ := ioutil.ReadAll(ctx.Request.Body)

		// TODO: we are not currently using the the request object. we should use it
		request := &messages.ThreadPageRequest{}
		if unmarshallErr := protojson.Unmarshal(body, request); unmarshallErr != nil {
			log.Println("error in marshalling: ", unmarshallErr)
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(unmarshallErr))
			return
		}
		var threads []entity.AgoraThread
		var displayThreads []*messages.ThreadDisplayInfo
		db.Order("created_at desc").Limit(10).Find(&threads)

		for _, thread := range threads {
			display, threadErr := createThreadDisplayInfo(db, thread.ID)
			if threadErr != nil {
				log.Println("cannot convert/find thread: ", thread.ID, thread.Title)
			}

			displayThreads = append(displayThreads, display)
		}
		ctx.JSON(http.StatusOK, displayThreads)
	}
}

func createThreadDisplayInfo(db *gorm.DB, id uint) (*messages.ThreadDisplayInfo, error) {
	var threadEntity entity.AgoraThread
	db.First(&threadEntity, "id = ? ", id)

	// TODO: error and nil checking here

	contributorsEntity, _ := queries.FindThreadContributors(db, threadEntity.ID)

	var contributorsProto []*messages.UserDisplayInfo
	var creatorProto *messages.UserDisplayInfo
	for _, contributor := range *contributorsEntity {
		log.Println(contributor.ContributorID)
		user := queries.FindAgoraUser(db, contributor.ContributorID)
		userDisplay := converters.ConvertUserToUserDisplayInfo(user, contributor.Access)

		if contributor.Access == messages.ContributeLevel_ThreadCreator {
			creatorProto = userDisplay
		} else {
			contributorsProto = append(contributorsProto, userDisplay)
		}
	}

	commentCount, _ := queries.CountThreadComments(db, threadEntity.ID) // TODO: nil and error checking

	return createInfoProto(threadEntity, creatorProto, contributorsProto, commentCount)
}

func createInfoProto(
	threadEntity entity.AgoraThread, creator *messages.UserDisplayInfo,
	contributors []*messages.UserDisplayInfo, comments int64) (*messages.ThreadDisplayInfo, error) {

	oneThread := messages.ThreadDisplayInfo{
		Title:              threadEntity.Title,
		Description:        threadEntity.Description,
		CreatedAt:          threadEntity.CreatedAt.UnixMilli(),
		ImageURL:           threadEntity.ImageURL,
		Creator:            creator,
		Contributors:       contributors,
		Watchers:           0,
		Comments:           comments,
		ThreadContributors: 0,
	}

	return &oneThread, nil

}

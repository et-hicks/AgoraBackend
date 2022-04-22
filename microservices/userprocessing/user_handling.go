package userprocessing

import (
	"errors"
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

func Login(db *gorm.DB) func(ctx *gin.Context) {
	log.Println("Logging in User")

	return func(ctx *gin.Context) {
		userInfo, marshalError := decomposeBody(ctx)
		if marshalError != nil {
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(marshalError))
			return
		}
		if nilChecking := nilCheck(userInfo, true); nilChecking != nil {
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(nilChecking))
			return
		}
		var existingUser *entity.AgoraUser
		result := db.First(existingUser, "username = ? or email = ?", userInfo.UserName, userInfo.Email) // DB Access
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(result.Error))
			return
		}
		if existingUser.Email == userInfo.Email {
			ctx.JSON(http.StatusOK, microserviceutils.GoodHTTP("user found", true))
		} else {
			ctx.JSON(http.StatusOK, microserviceutils.GoodHTTP("no user found", false))
		}

	}
}

func CreateUser(db *gorm.DB) func(ctx *gin.Context) {
	log.Println("Creating new User")

	return func(ctx *gin.Context) {
		userInfo, marshalError := decomposeBody(ctx)
		if marshalError != nil {
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(marshalError))
			return
		}
		if nilChecking := nilCheck(userInfo, false); nilChecking != nil {
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(nilChecking))
			return
		}
		result := db.Create(CreateUserEntity(userInfo)) // DB Access
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(result.Error))
			return
		}
		ctx.JSON(http.StatusOK, microserviceutils.GoodHTTP("User Created Successfully", true))
	}
}

func CreateUserEntity(userInfo *messages.UserInfo) *entity.AgoraUser {

	return &entity.AgoraUser{
		FirstName:       userInfo.DisplayFirstName,
		LastName:        userInfo.DisplayLastName,
		Username:        userInfo.UserName,
		Email:           userInfo.Email,
		Password:        userInfo.Password,
		PhoneNumber:     userInfo.PhoneNumber,
		PhoneCode:       userInfo.PhoneCode,
		Type:            userInfo.Type,
	}
}


func nilCheck(userInfo *messages.UserInfo, forLogin bool) error {

	if userInfo.UserName == "" {
		return errors.New("no username set")
	}
	if userInfo.Password == "" {
		return errors.New("no password found")
	}

	if forLogin {
		return nil
	}

	// Since the user display name is optional, I am now only checking for email
	// for now
	if userInfo.Email == "" {
		return errors.New("no email found for create")
	}

	return nil
}

func decomposeBody(ctx *gin.Context) (*messages.UserInfo, error) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	var userInfo *messages.UserInfo
	if unmarshallErr := protojson.Unmarshal(body, userInfo); unmarshallErr != nil {
		return nil, unmarshallErr
	}
	return userInfo, nil
}

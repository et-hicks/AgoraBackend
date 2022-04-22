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
		existingUser := &entity.AgoraUser{}
		result := db.First(existingUser, "username = ? or email = ?", userInfo.Username, userInfo.Email) // DB Access

		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(result.Error))
			return
		}
		if existingUser.ID == 0 {
			ctx.JSON(http.StatusOK, microserviceutils.GoodHTTP("no user found", false))
		}
		message, match, foundError := found(userInfo, existingUser)
		if foundError != nil {
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(foundError))
			return
		}
		ctx.JSON(http.StatusOK, microserviceutils.GoodHTTP(message, match))
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
		// TODO: validate email string

		// Check if they exist already
		existingUser := &entity.AgoraUser{}
		db.First(existingUser, "username = ? or email = ?", userInfo.Username, userInfo.Email) // DB Access
		if existingUser.ID != 0 {
			ctx.JSON(http.StatusOK, microserviceutils.GoodHTTP("username/email already exists", false))
			return
		}

		result := db.Create(UserEntityFactory(userInfo)) // DB Access
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, microserviceutils.BadHTTP(result.Error))
			return
		}
		ctx.JSON(http.StatusOK, microserviceutils.GoodHTTP("User Created Successfully", true))
	}
}

func found(userInfo *messages.UserInfo, userEntity *entity.AgoraUser) (string, bool, error) {
	password, err := microserviceutils.Encrypt(userInfo.Password, microserviceutils.MySecret)
	if err != nil {
		return "", false, err
	}
	passwordMatch := password == userEntity.Password
	message := "No user found"

	if userInfo.Username == "" {
		emailsMatch := userEntity.Email == userInfo.Email
		match := passwordMatch && emailsMatch
		if match {
			message = "user successfully found"
		}
		return message, passwordMatch && emailsMatch, nil
	}
	usernameMatch := userInfo.Username == userEntity.Username
	match := passwordMatch && usernameMatch
	if match {
		message = "user successfully found"
	}
	return message, match, nil
}

func UserEntityFactory(userInfo *messages.UserInfo) *entity.AgoraUser {

	password, err := microserviceutils.Encrypt(userInfo.Password, microserviceutils.MySecret)
	if err != nil {
		return nil
	}
	return &entity.AgoraUser{
		FirstName:       userInfo.DisplayFirstName,
		LastName:        userInfo.DisplayLastName,
		Username:        userInfo.Username,
		Email:           userInfo.Email,
		Password:        password,
		PhoneNumber:     userInfo.PhoneNumber,
		PhoneCode:       userInfo.PhoneCode,
		Type:            userInfo.Type,
	}
}


func nilCheck(userInfo *messages.UserInfo, forLogin bool) error {

	if userInfo.Username == "" && userInfo.Email == "" {
		return errors.New("no username/email set")
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
	userInfo := &messages.UserInfo{}
	if unmarshallErr := protojson.Unmarshal(body, userInfo); unmarshallErr != nil {
		return nil, unmarshallErr
	}
	return userInfo, nil
}

package microserviceutils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type HTTPResponse struct {
	Message string `json:"message"`
	Success bool `json:"success"`
}

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// MySecret This should be in an env file in production
const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

func GeneralInit() (*gorm.DB, *gin.Engine, error) {
	dosn := "agora_mysql_admin:DigitalMcDonalds$3.21@tcp(agora-mysql-dev.mysql.database.azure.com)/agora?charset=utf8mb4&parseTime=true"
	fmt.Println(dosn)
	gormDb, sqlErr := gorm.Open(mysql.Open(dosn), &gorm.Config{})

	// and the router
	router := gin.Default()

	if sqlErr != nil {
		return nil, nil, sqlErr
	}

	return gormDb, router, nil
}

func HardReset(*gin.Context) {
	// TODO: Learn how to gracefully shut down/restart the server
	log.Println("Ungraceful shutdown of the server")
	log.Fatal("Resetting server")
}

func BadHTTP(err error) HTTPResponse {
	return HTTPResponse{Message: err.Error()}
}

func GoodHTTP(response string, success bool) HTTPResponse {
	return HTTPResponse{Message: response, Success: success}
}


// thanks to Log Rocket
// https://blog.logrocket.com/learn-golang-encryption-decryption/

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Encrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func NotImplementedYet() error {
	return errors.New("method is not implemented yet")
}




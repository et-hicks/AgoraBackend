package main

import (
	"fmt"
	"github.com/admin-agora/backend/messages"
	"github.com/admin-agora/backend/microservices/commentprocessing"
	"github.com/admin-agora/backend/microservices/threadprocessing"
	"github.com/admin-agora/backend/microservices/userprocessing"
	"github.com/admin-agora/backend/sql/entity"
	_ "github.com/denisenkom/go-mssqldb"
	"os"

	// "gorm.io/driver/sqlserver"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)


type Employees struct {
	Id uint
	Name string
	Location string
}

func gormExample() {

	//gormDb, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	dosn := "agora_mysql_admin:DigitalMcDonalds$3.21@tcp(agora-mysql-dev.mysql.database.azure.com)/agora?charset=utf8mb4"
	fmt.Println(dosn)
	gormDb, err := gorm.Open(mysql.Open(dosn), &gorm.Config{})
	if err != nil {
		fmt.Println("fuck")
		return
	}

	var newE = Employees{
		2,
		"Ethan",
		"New york",
	}
	gormDb.Create(&newE)
	if err != nil {
		log.Fatal("cannot validate connection")
	}
	var employee Employees
	// gormDb.First(&employee)
	fmt.Println(employee)
	fmt.Println("i honestly cannot believe this fucking worked")
}

func BuildUser(user *entity.AgoraUser) {

	user.Username = "etcpie"
	user.FirstName = "Ethan"
	user.LastName = "Hicks"
	user.Email = "etmhicks@gmail.com"
	user.Password = "DigitalMcDonalds&3.21"
	user.Type = messages.AccountType(4)
	user.BEFSJson =`{
	   "name":"John",
	   "age":29,
	   "hobbies":[
		  "martial arts",
		  "breakfast foods",
		  "piano"
	   ]
	}`
	// user.BEFS := make(map[string]interface{})


}

func main() {

	// TODO: set the IP address for the azure database to be valid in cloud run

	name := os.Getenv("K_SERVICE")
	switch name {
	case "agoracomments":
		commentprocessing.Service()
	case "agorathreads":
		threadprocessing.Service()
	case "agorauser":
		userprocessing.Service()
	default:
		log.Println("nothing found for name")
	}
	threadprocessing.Service()
	//commentprocessing.Service()
	// gormTest()
	//commentprocessing.Service()
	//threadprocessing.Service()
	//userprocessing.Service()

	//SQLSetUp(nil, nil)
}

func gormTest() {
	// Create connection pool
	gormExample()

	dosn := "agora_mysql_admin:DigitalMcDonalds$3.21@tcp(agora-mysql-dev.mysql.database.azure.com)/agora?charset=utf8mb4&parseTime=true"
	fmt.Println(dosn)
	gormDb, sqlErr := gorm.Open(mysql.Open(dosn), &gorm.Config{})

	if sqlErr != nil {
		fmt.Println("fuck, but in connection")
	}

	var user entity.AgoraUser

	BuildUser(&user)
	fmt.Println("pause here please and thank your")
	gormDb.Create(&user)

	var ethan entity.AgoraUser
	gormDb.First(&ethan)
	errE := ethan.LoadForCode()
	if errE != nil {
		return
	}

	fmt.Println("pause here please and thank your 2")

	newFEFS := make(map[string]interface{})
	newFEFS["frontEndName"] = "Johnny Sins"
	newFEFS["backEndName"] = "Anal McGee"
	ethan.FEFS = newFEFS
	ethanUnloadError := ethan.UnloadForDatabase()
	if ethanUnloadError != nil {
		fmt.Println("fuck")
	}
	gormDb.Updates(ethan)

}


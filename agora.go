package main

import (
	"fmt"
	"github.com/admin-agora/backend/sql/entity"
	_ "github.com/denisenkom/go-mssqldb"
	"net/http"

	// "gorm.io/driver/sqlserver"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)


type Employees struct {
	Id uint64
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

func BuildUser(user *entity.User) {

	user.Username = "etcpie"
	user.FirstName = "Ethan"
	user.LastName = "Hicks"
	user.Email = "etmhicks@gmail.com"
	user.Password = "DigitalMcDonalds&3.21"
	user.Type = entity.Employee
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

	http.HandleFunc("/", CommentPostService)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}
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

	var user entity.User

	BuildUser(&user)
	fmt.Println("pause here please and thank your")
	gormDb.Create(&user)

	var ethan entity.User
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


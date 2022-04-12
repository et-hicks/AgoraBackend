package entity

import (
	"encoding/json"
	"gorm.io/gorm"
)

type AccountType int64
const (
	Contributor AccountType = iota  // thread creator
	Commentary						// can make comments but not threads
	Login 							// login and just view stuff
	Employee						// Employees of Agora
)



type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Username    string
	Email       string
	Password    string
	PhoneNumber uint64
	PhoneCode   uint64
	Type     AccountType
	FunctionalStuff
}

func (u *User) LoadForCode() error {
	befs := make(map[string]interface{})
	fefs := make(map[string]interface{})

	if u.BEFSJson == "" {
		u.BEFSJson = "{}"
	}
	if u.FEFSJson == "" {
		u.FEFSJson = "{}"
	}

	// TODO: handle null values here
	errFefs := json.Unmarshal([]byte(u.BEFSJson), &befs)
	errBefs := json.Unmarshal([]byte(u.FEFSJson), &fefs)
	if errFefs != nil {
		return errFefs
	}
	if errBefs != nil {
		return errBefs
	}

	u.FEFS = fefs
	u.BEFS = befs

	return nil
}

func (u *User) UnloadForDatabase() error {
	fefsJson, errFefs := json.Marshal(u.FEFS)
	befsJson, errBefs := json.Marshal(u.BEFS)

	if errFefs != nil {
		return errFefs
	}
	if errBefs != nil {
		return errBefs
	}

	u.FEFSJson = string(fefsJson)
	u.BEFSJson = string(befsJson)

	return nil
}


func (u *User) CreateTable(db *gorm.DB) error {
	createTableSql := `
		CREATE TABLE users (
		id bigint NOT NULL AUTO_INCREMENT,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL,
		deleted_at CHAR(255) CHARACTER SET UTF8MB4,
		first_name CHAR(255) CHARACTER SET UTF8MB4 DEFAULT NULL,
		last_name CHAR(255) CHARACTER SET UTF8MB4 DEFAULT NULL,
		username CHAR(255) CHARACTER SET UTF8MB4 NOT NULL,
		email CHAR(255) CHARACTER SET UTF8MB4 NOT NULL,
		password CHAR(255) CHARACTER SET UTF8MB4 DEFAULT NULL,
		phone_number CHAR(255) CHARACTER SET UTF8MB4 DEFAULT NULL,
    	phone_code CHAR(255) CHARACTER SET UTF8MB4 DEFAULT NULL,
		type	CHAR(255) CHARACTER SET UTF8MB4 DEFAULT NULL,
		fefs_json TEXT DEFAULT NULL,
		befs_json TEXT DEFAULT NULL,
		
		Primary Key(ID)
	) ENGINE=InnoDB;
	`

	db.Exec(createTableSql)

	return nil
}

func (u *User) UpdateTable(db *gorm.DB, sqlUpdate string) error {
	db.Exec(sqlUpdate)

	return nil
}


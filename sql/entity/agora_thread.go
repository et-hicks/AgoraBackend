package entity

import (
	"encoding/json"
	"github.com/admin-agora/backend/messages"
	"gorm.io/gorm"
)

// TODO: Determine if I want comments in the thread or not
// for now, not in the thread entity

type AgoraThread struct {
	gorm.Model
	Title string
	Description string
	CreatorID uint
	Likes uint64
	Dislikes uint64
	Clicks uint64
	Watchers uint64
	UrlUUID string
	ImageURL string
	Public messages.PublicityLevel
	IsReported bool
	FunctionalStuff
}

func (t *AgoraThread) LoadForCode() error {
	befs := make(map[string]interface{})
	fefs := make(map[string]interface{})

	if t.BEFSJson == "" {
		t.BEFSJson = "{}"
	}
	if t.FEFSJson == "" {
		t.FEFSJson = "{}"
	}

	// TODO: handle null values here
	errFefs := json.Unmarshal([]byte(t.BEFSJson), &befs)
	errBefs := json.Unmarshal([]byte(t.FEFSJson), &fefs)
	if errFefs != nil {
		return errFefs
	}
	if errBefs != nil {
		return errBefs
	}

	t.FEFS = fefs
	t.BEFS = befs

	return nil
}

func (t *AgoraThread) UnloadForDatabase() error {
	fefsJson, errFefs := json.Marshal(t.FEFS)
	befsJson, errBefs := json.Marshal(t.BEFS)

	if errFefs != nil {
		return errFefs
	}
	if errBefs != nil {
		return errBefs
	}

	t.FEFSJson = string(fefsJson)
	t.BEFSJson = string(befsJson)

	return nil
}

func (t *AgoraThread) CreateTable(db *gorm.DB) error {
	createTableSql := `
		CREATE TABLE IF NOT EXISTS agora_threads (
			id bigint NOT NULL AUTO_INCREMENT,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			deleted_at CHAR(255) CHARACTER SET UTF8MB4 DEFAULT NULL,

			title CHAR(255) CHARACTER SET UTF8MB4 NOT NULL,
			description CHAR(255) CHARACTER SET UTF8MB4 DEFAULT NULL,
			creator_id BIGINT NOT NULL,
			likes BIGINT DEFAULT NULL,
			dislikes BIGINT DEFAULT NULL,
			clicks BIGINT DEFAULT NULL,
			watchers BIGINT DEFAULT NULL,
			url_uuid CHAR(255) CHARACTER SET UTF8MB4 NOT NULL,
			image_url CHAR(255) DEFAULT NULL,
			public CHAR(255) CHARACTER SET UTF8MB4 NOT NULL,
			is_reported INT NOT NULL,

			fefs_json TEXT DEFAULT NULL,
			befs_json TEXT DEFAULT NULL,
			
			PRIMARY KEY(id),
			CONSTRAINT fk_id_usersId FOREIGN KEY(creator_id) REFERENCES users(id)
	
		) ENGINE=InnoDB;
	`


	db.Exec(createTableSql)

	return nil
}

func (t *AgoraThread) UpdateTable(db *gorm.DB, sqlUpdate string) error {
	db.Exec(sqlUpdate)

	return nil
}
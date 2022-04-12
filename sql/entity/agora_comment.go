package entity

import (
	"encoding/json"
	"gorm.io/gorm"
)

type AgoraComment struct {
	gorm.Model
	AuthorID User `gorm:"foreignKey:ID"`
	ParentCommentId uint64 // TODO: learn how to deal with this
	Comment 		string
	IsEdited		bool
	IsReported 		bool
	Likes			uint64
	Dislikes		uint64
	FunctionalStuff
}

func (c *AgoraComment) LoadForCode() error {
	befs := make(map[string]interface{})
	fefs := make(map[string]interface{})

	if c.BEFSJson == "" {
		c.BEFSJson = "{}"
	}
	if c.FEFSJson == "" {
		c.FEFSJson = "{}"
	}

	// TODO: handle null values here
	errFefs := json.Unmarshal([]byte(c.BEFSJson), &befs)
	errBefs := json.Unmarshal([]byte(c.FEFSJson), &fefs)
	if errFefs != nil {
		return errFefs
	}
	if errBefs != nil {
		return errBefs
	}

	c.FEFS = fefs
	c.BEFS = befs

	return nil
}

func (c *AgoraComment) UnloadForDatabase() error {
	fefsJson, errFefs := json.Marshal(c.FEFS)
	befsJson, errBefs := json.Marshal(c.BEFS)

	if errFefs != nil {
		return errFefs
	}
	if errBefs != nil {
		return errBefs
	}

	c.FEFSJson = string(fefsJson)
	c.BEFSJson = string(befsJson)

	return nil
}

func (c *AgoraComment) CreateTable(db *gorm.DB) error {
	createTableSql := `
		CREATE TABLE agora_comments (
		id bigint NOT NULL AUTO_INCREMENT,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL,
		deleted_at CHAR(255) CHARACTER SET UTF8MB4,
		
        author_id BIGINT,
        parent_comment_id BIGINT DEFAULT NULL,
		comment TEXT DEFAULT NULL,
		is_edited INT NOT NULL,
        is_reported INT NOT NULL,
        likes BIGINT DEFAULT NULL,
		dislikes BIGINT DEFAULT NULL,
		
		fefs_json TEXT DEFAULT NULL,
		befs_json TEXT DEFAULT NULL,
		
		PRIMARY KEY(id),
        CONSTRAINT fk_comments_id_usersId FOREIGN KEY(author_id) REFERENCES users(id)

	) ENGINE=InnoDB;
	`


	db.Exec(createTableSql)

	return nil
}

func (c *AgoraComment) UpdateTable(db *gorm.DB, sqlUpdate string) error {
	db.Exec(sqlUpdate)

	return nil
}

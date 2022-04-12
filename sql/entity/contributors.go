package entity

import (
	"gorm.io/gorm"
)

type AccessLevel int64
const (
	Undefined AccessLevel = iota
	Admin
	Moderator
	Creator
	Commenter
	Revoked
	Blocked
	Viewer
)

type Contribute struct {
	gorm.Model
	Contributor User      	`gorm:"foreignKey:ID"`
	Thread		AgoraThread `gorm:"foreignKey:ID"`
	Access AccessLevel       // to define how they can interact with the thread
}

func (c *Contribute) LoadForCode() error {

	return nil
}

func (c *Contribute) UnloadForDatabase() error {

	return nil
}


func (c *Contribute) CreateTable(db *gorm.DB) error {
	createTableSql := `
		CREATE TABLE contributors (
			id bigint NOT NULL AUTO_INCREMENT,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			deleted_at CHAR(255) CHARACTER SET UTF8MB4,
		
			contributor_id BIGINT,
			thread_id BIGINT,
			access CHAR(255) CHARACTER SET UTF8MB4 NOT NULL,
		
			PRIMARY KEY(id),
			CONSTRAINT fk_contributor_id_users_id FOREIGN KEY(contributor_id) REFERENCES users(id),
			CONSTRAINT fk_contributor_id_threads_id FOREIGN KEY(thread_id) REFERENCES threads(id)

		) ENGINE=InnoDB;
	`


	db.Exec(createTableSql)

	return nil
}

func (c *Contribute) UpdateTable(db *gorm.DB, sqlUpdate string) error {
	db.Exec(sqlUpdate)

	return nil
}

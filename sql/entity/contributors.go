package entity

import (
	"github.com/admin-agora/backend/messages"
	"gorm.io/gorm"
)



type Contributor struct {
	gorm.Model
	ContributorID uint
	ThreadID      uint
	Access        messages.ContributeLevel
}

func (c *Contributor) LoadForCode() error {

	return nil
}

func (c *Contributor) UnloadForDatabase() error {

	return nil
}


func (c *Contributor) CreateTable(db *gorm.DB) error {
	createTableSql := `
		CREATE TABLE contributors (
			id bigint NOT NULL AUTO_INCREMENT,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			deleted_at CHAR(255) CHARACTER SET UTF8MB4,
		
			contributor_id BIGINT,
			thread_id BIGINT,
			access CHAR(255) CHARACTER SET UTF8MB4 NOT NULL,
		
			PRIMARY KEY(id)

		) ENGINE=InnoDB;
	`


	db.Exec(createTableSql)

	return nil
}

func (c *Contributor) AddConstraints(db *gorm.DB) error {
	// i dont think gorm requires semicolons at the end
	// TODO: determine if gorm needs semicolons
	threadsConstraint :=
		"alter table contributors add CONSTRAINT if not exists fk_contributor_id_threads_id FOREIGN KEY(thread_id) REFERENCES threads(id)"
	usersConstraint :=
		"alter table contributors add CONSTRAINT if not exists fk_contributor_id_users_id FOREIGN KEY(contributor_id) REFERENCES users(id)"

	// TODO: return errors if exist
	db.Exec(threadsConstraint)
	db.Exec(usersConstraint)
	return nil
}

func (c *Contributor) DeleteConstraints(db *gorm.DB) error {
	// i dont think gorm requires semicolons at the end
	// TODO: determine if gorm needs semicolons
	threadsConstraint :=
		"alter table thread_comments DROP FOREIGN KEY if exists fk_contributor_id_threads_id"
	usersConstraint :=
		"alter table thread_comments DROP FOREIGN KEY if exists fk_contributor_id_users_id"

	// TODO: return errors if exist
	db.Exec(threadsConstraint)
	db.Exec(usersConstraint)
	return nil
}

func (c *Contributor) UpdateTable(db *gorm.DB, sqlUpdate string) error {
	db.Exec(sqlUpdate)

	return nil
}

package entity

import "gorm.io/gorm"

type TopicCategory int64
const (
	Hashtag TopicCategory = iota
	Category
)

type Topic struct {
	gorm.Model
	Thread		AgoraThread `gorm:"foreignKey:ID"`
	Topic string
	Category TopicCategory
}

func (t *Topic) LoadForCode() error {

	return nil
}

func (t *Topic) UnloadForDatabase() error {

	return nil
}

func (t *Topic) CreateTable(db *gorm.DB) error {
	createTableSql := `
		CREATE TABLE topics (
			id bigint NOT NULL AUTO_INCREMENT,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			deleted_at CHAR(255) CHARACTER SET UTF8MB4,
		
			thread_id BIGINT,
			topic CHAR(255) CHARACTER SET UTF8MB4 NOT NULL,
			topic_category CHAR(16) CHARACTER SET UTF8MB4 NOT NULL,
		
			PRIMARY KEY(id),
			CONSTRAINT fk_topic_id_threads_id FOREIGN KEY(thread_id) REFERENCES threads(id)
		
		) ENGINE=InnoDB;
	`


	db.Exec(createTableSql)

	return nil
}

func (t *Topic) UpdateTable(db *gorm.DB, sqlUpdate string) error {
	db.Exec(sqlUpdate)

	return nil
}

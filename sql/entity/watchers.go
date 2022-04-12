package entity

import "gorm.io/gorm"

type WatcherStatus int64
const (
	Notification WatcherStatus = iota
	Email
	Summary      // After the thread has died down, notify top X comments
	EmailSummary // After the thread has died down, notify and Email top X comments
)
type Watcher struct {
	gorm.Model
	Watcher User      	`gorm:"foreignKey:ID"`
	Thread		AgoraThread `gorm:"foreignKey:ID"`
	Status WatcherStatus
}
func (w *Watcher) LoadForCode() error {

	return nil
}

func (w *Watcher) UnloadForDatabase() error {

	return nil
}

func (w *Watcher) CreateTable(db *gorm.DB) error {
	createTableSql := `
		CREATE TABLE watchers (
			id bigint NOT NULL AUTO_INCREMENT,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			deleted_at CHAR(255) CHARACTER SET UTF8MB4,
		
			watcher_id BIGINT,
			thread_id BIGINT,
			access CHAR(255) CHARACTER SET UTF8MB4 NOT NULL,
		
			PRIMARY KEY(id),
			CONSTRAINT fk_watcher_id_users_id FOREIGN KEY(watcher_id) REFERENCES users(id),
			CONSTRAINT fk_watcher_id_threads_id FOREIGN KEY(thread_id) REFERENCES threads(id)

		) ENGINE=InnoDB;
	`


	db.Exec(createTableSql)

	return nil
}

func (w *Watcher) UpdateTable(db *gorm.DB, sqlUpdate string) error {
	db.Exec(sqlUpdate)

	return nil
}

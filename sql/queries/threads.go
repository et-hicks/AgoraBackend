package queries

import (
	"github.com/admin-agora/backend/sql/entity"
	"github.com/admin-agora/backend/src/utils/microserviceutils"
	"gorm.io/gorm"
)

func CountThreadComments(db *gorm.DB, threadID uint) (int64, error) {
	var count int64
	countComments :=
		"select count(*) from agora_comments where thread_id = ?"
	db.Raw(countComments, threadID).Scan(&count)
	return count, nil
}

func countThreadWatchers(db *gorm.DB, threadID string) (int64, error) {
	// TODO: impl later
	return 0, microserviceutils.NotImplementedYet()
}

func FindThreadContributors(db *gorm.DB, threadID uint) (*[]entity.Contributor, error) {

	var contributors []entity.Contributor

	result := db.Where("thread_id = ?", threadID).Find(&contributors) // DB Access

	if result.Error != nil {
		return nil, result.Error
	}

	return &contributors, nil
}

func FindThreadComments(db *gorm.DB, threadID uint) (*[]entity.AgoraComment, error) {

	// TODO: somehow enforce some ordering here

	var comments []entity.AgoraComment

	result := db.Where("thread_id = ?", threadID).Find(&comments) // DB Access

	if result.Error != nil {
		return nil, result.Error
	}

	return &comments, nil
}


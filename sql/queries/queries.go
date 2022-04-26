package queries

import (
	"github.com/admin-agora/backend/sql/entity"
	"github.com/admin-agora/backend/src/utils/microserviceutils"
	"gorm.io/gorm"
)

func countThreadComments(db *gorm.DB, threadId string) (int64, error) {
	var count int64
	countComments :=
		"select count(*) from comments where thread_id = ?"
	db.Raw(countComments, threadId).Scan(&count)
	return count, nil
}

func countThreadWatchers(db *gorm.DB, threadId string) (int64, error) {
	// TODO: impl later
	return 0, microserviceutils.NotImplementedYet()
}

func findThreadContributors(db *gorm.DB, threadID string) ([]*entity.AgoraUser, error) {

	var contributors []*entity.AgoraUser

	// IDK how gorm is supposed to work here. Imma do what I know how to do
	findContributors :=
		"select * from agora_users " +
			"left join contributors on contributors.contributor_id = agora_users.id " +
			"where thread_id = ?"

	rows, err := db.Raw(findContributors).Scan().Rows()
	if err != nil {
		return nil, err
	}

}




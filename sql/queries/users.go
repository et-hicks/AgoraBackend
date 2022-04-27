package queries

import (
	"github.com/admin-agora/backend/sql/entity"
	"gorm.io/gorm"
)

// Using gorm properly will mean this is not required
// but this is faster for now than the other method soooo....

func FindAgoraUser(db *gorm.DB, userID uint) *entity.AgoraUser {
	var existingUser entity.AgoraUser
	db.First(&existingUser, "id = ?", userID) // DB Access
	return &existingUser
}




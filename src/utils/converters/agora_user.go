package converters

import (
	"github.com/admin-agora/backend/messages"
	"github.com/admin-agora/backend/sql/entity"
)

func ConvertUserToUserDisplayInfo(user *entity.AgoraUser, level messages.ContributeLevel) *messages.UserDisplayInfo {

	userDisplay := messages.UserDisplayInfo{
		Username:     user.Username,
		UserPageURL:  "/path/to/users/page",
		UserImageURL: user.ImageURL,
		ContributeLevel:        level,
	}
	return &userDisplay
}

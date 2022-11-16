package followers

import (
	"backend/pkg/chat"
	"backend/pkg/follow"
	"backend/pkg/structs"
)
//Get will return a slice of all users following userId
func Get(UserId string, database structs.DB) ([]structs.Info, error) {
	var userInfos []structs.Info
	followers, err := follow.GetFollowers(UserId, &database)
	if err != nil {
		return nil, err
	}
	for _, follower := range followers {
		userInfo, err := chat.GetUserInfo(follower.FollowerId, &database)
		if err != nil {
			return nil, err
		}
		userInfos = append(userInfos, userInfo)
	}
	return userInfos, nil
}

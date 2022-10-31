package follow

import (
	"time"

	"backend/pkg/helper"
	l "backend/pkg/log"
	"backend/pkg/structs"
)

// InsertFollow will insert a follow into the database
//
// Params:
//
//	followerId: the id of the user who is following - current user
//	followingId: the id of the user who is being followed - other user
//	database: the database to insert the follow into
func InsertFollow(follwerId, followingId string, database *structs.DB) error {
	createdAt := time.Now().String()
	stmt, err := database.DB.Prepare("INSERT INTO Follower (followerId, followingId, createdAt) VALUES (?, ?, ?)")
	if err != nil {
		l.LogMessage("follow.go", "InsertFollow", err)
		return err
	}
	_, err = stmt.Exec(follwerId, followingId, createdAt)
	if err != nil {
		l.LogMessage("follow.go", "InsertFollow", err)
		return err
	}
	return nil
}

// DeleteFollow will delete a follow from the database
//
// Params:
//
//	followerId: the id of the user who is following
//	followingId: the id of the user who is being followed
//	database: the database to delete the follow from
func DeleteFollow(followerId, followingId string, database *structs.DB) error {
	stmt, err := database.DB.Prepare("DELETE FROM Follower WHERE followerId = ? AND followingId = ?")
	if err != nil {
		l.LogMessage("follow.go", "DeleteFollow", err)
		return err
	}
	_, err = stmt.Exec(followerId, followingId)
	if err != nil {
		l.LogMessage("follow.go", "DeleteFollow", err)
		return err
	}
	stmt, err = database.DB.Prepare("DELETE FROM FollowNotif WHERE userId = ? AND followingId = ?")
	if err != nil {
		l.LogMessage("follow.go", "DeleteFollow", err)
		return err
	}
	_, err = stmt.Exec(followerId, followingId)
	if err != nil {
		l.LogMessage("follow.go", "DeleteFollow", err)
		return err
	}
	return nil
}

// CheckIfFollow checks if current user is following the other user
//
// Params:
//
//	followerId: the id of the user who is following - current user
//	followingId: the id of the user who is being followed - other user
//	database: the database to check the follow in
func CheckIfFollow(followerId, followingId string, database *structs.DB) bool {
	// check if a user is following another user
	stmt, err := database.DB.Query("SELECT * FROM Follower WHERE followerId = ? AND followingId = ?", followerId, followingId)
	if err != nil {
		l.LogMessage("follow.go", "CheckFollow", err)
		return false
	}
	var follower structs.Follower
	for stmt.Next() {
		stmt.Scan(
			&follower.FollowerId,
			&follower.FollowingId,
			&follower.CreatedAt,
		)
	}
	return follower != structs.Follower{}
}

// InsertFollowNotif will insert a follow notification into the database
//
// Params:
//
//	followerId: the id of the user who is following - current user
//	followingId: the id of the user who is being followed - other user
//	status: the status of the follow (follow or pending)
//	database: the database to insert the follow notification into
func InsertFollowNotif(followerId, followingId, status string, database *structs.DB) error {
	createdAt := time.Now().String()
	stmt, err := database.DB.Prepare("INSERT INTO FollowNotif (userId, followingId, status, createdAt, unread) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		l.LogMessage("follow.go", "InsertFollowNotif", err)
		return err
	}
	_, err = stmt.Exec(followerId, followingId, status, createdAt, true)
	if err != nil {
		l.LogMessage("follow.go", "InsertFollowNotif", err)
		return err
	}
	return nil
}

// GetNumOfFollowing will get the number of users the current user is following
//
// Params:
//
//	userId: the id of the user who is following - current user
//	database: the database to get the number of following from
func GetNumOfFollowing(userId string, database *structs.DB) (int, error) {
	stmt, err := database.DB.Query("SELECT numFollowing FROM User WHERE userId = ?", userId)
	if err != nil {
		l.LogMessage("follow.go", "GetNumOfFollowing", err)
		return 0, nil
	}
	var user structs.User
	for stmt.Next() {
		stmt.Scan(
			&user.NumFollowing,
		)
	}
	l.LogMessage("follow.go", "GetNumOfFollowing", user)
	return user.NumFollowing, nil
}

// GetNumOfFollowers will get the number of users following the current user
//
// Params:
//
//	userId: the id of the user who is being followed - current user
//	database: the database to get the number of followers from
func GetNumOfFollowers(userId string, database *structs.DB) (int, error) {
	stmt, err := database.DB.Query("SELECT numFollowers FROM User WHERE userId = ?", userId)
	if err != nil {
		l.LogMessage("follow.go", "GetNumOfFollowers", err)
		return 0, nil
	}
	var user structs.User
	for stmt.Next() {
		stmt.Scan(
			&user.NumFollowers,
		)
	}
	l.LogMessage("follow.go", "GetNumOfFollowers", user)
	return user.NumFollowers, nil
}

// UpdateNumOfFollowing will update the number of users the current user is following
//
// Params:
//
//	userId: the id of the user who is following - current user
//	numFollowing: the number of users the current user is following
//	database: the database to update the number of following in
func UpdateNumOfFollowing(userId string, numFollowing int, database *structs.DB) error {
	stmt, err := database.DB.Prepare("UPDATE User SET numFollowing = ? WHERE userId = ?")
	if err != nil {
		l.LogMessage("follow.go", "UpdateNumOfFollowing", err)
		return err
	}
	_, err = stmt.Exec(numFollowing, userId)
	if err != nil {
		l.LogMessage("follow.go", "UpdateNumOfFollowing", err)
		return err
	}
	return nil
}

// UpdateNumOfFollowers will update the number of users following the current user
//
// Params:
//
//	userId: the id of the user who is being followed - current user
//	numFollowers: the number of users following the current user
//	database: the database to update the number of followers in
func UpdateNumOfFollowers(userId string, numFollowers int, database *structs.DB) error {
	stmt, err := database.DB.Prepare("UPDATE User SET numFollowers = ? WHERE userId = ?")
	if err != nil {
		l.LogMessage("follow.go", "UpdateNumOfFollowers", err)
		return err
	}
	_, err = stmt.Exec(numFollowers, userId)
	if err != nil {
		l.LogMessage("follow.go", "UpdateNumOfFollowers", err)
		return err
	}
	return nil
}

// UpdateFollowNotifStatus will update the status of a follow notification
//
// Params:
//
//	followerId: the id of the user who is following - current user
//	followingId: the id of the user who is being followed - other user
//	status: the status of the follow (follow or pending)
//	database: the database to update the follow notification status in
func UpdateFollowNotifStatus(followerId, followingId, status string, database *structs.DB) error {
	stmt, err := database.DB.Prepare("UPDATE FollowNotif SET status = ? WHERE userId = ? AND followingId = ?")
	if err != nil {
		l.LogMessage("follow.go", "UpdateFollowNotifStatus", err)
		return err
	}
	_, err = stmt.Exec(status, followerId, followingId)
	if err != nil {
		l.LogMessage("follow.go", "UpdateFollowNotifStatus", err)
		return err
	}
	return nil
}

// FollowUser will follow a user if the user does not already follow the other user
// and delete the follow if the user does follow the other user
//
//	returns status (follow, pending or  unfollow) and error
//
// Params:
//
//	followerId: the id of the user who is following - current user
//	followingId: the id of the user who is being followed - other user
//	database: the database to insert the follow into
func FollowUser(followerId, followingId string, database *structs.DB) (string, error) {
	if CheckIfFollow(followerId, followingId, database) {
		DeleteFollow(followerId, followingId, database)
		currentUserNumOfFollowing, _ := GetNumOfFollowing(followerId, database)
		UpdateNumOfFollowing(followerId, currentUserNumOfFollowing-1, database)
		otherUserNumOfFollowers, _ := GetNumOfFollowers(followingId, database)
		UpdateNumOfFollowers(followingId, otherUserNumOfFollowers-1, database)
		return "unfollow", nil
	}
	if helper.CheckUserIfPublic(followingId, database) {
		InsertFollow(followerId, followingId, database)
		currentUserNumOfFollowing, _ := GetNumOfFollowing(followerId, database)
		UpdateNumOfFollowing(followerId, currentUserNumOfFollowing+1, database)
		otherUserNumOfFollowers, _ := GetNumOfFollowers(followingId, database)
		UpdateNumOfFollowers(followingId, otherUserNumOfFollowers+1, database)
		InsertFollowNotif(followerId, followingId, "follow", database)
		return "follow", nil
	}
	InsertFollowNotif(followerId, followingId, "pending", database)
	return "pending", nil
}

// AcceptFollow will accept a follow request from another user
//
// Params:
//
//	followerId: the id of the user who is following - other user
//	followingId: the id of the user who is being followed - current user
//	database: the database to insert the follow into
func AcceptFollow(followerId, followingId string, database *structs.DB) {
	InsertFollow(followerId, followingId, database)
	currentUserNumOfFollowing, _ := GetNumOfFollowing(followerId, database)
	UpdateNumOfFollowing(followerId, currentUserNumOfFollowing+1, database)
	otherUserNumOfFollowers, _ := GetNumOfFollowers(followingId, database)
	UpdateNumOfFollowers(followingId, otherUserNumOfFollowers+1, database)
	UpdateFollowNotifStatus(followerId, followingId, "follow", database)
}

// DeclineFollow will decline a follow request from another user
//
// Params:
//
//	followerId: the id of the user who is following - other user
//	followingId: the id of the user who is being followed - current user
//	database: the database to insert the follow into
func DeclineFollow(followerId, followingId string, database *structs.DB) {
	DeleteFollow(followerId, followingId, database)
}

package models

import "time"

// Member model
type Member struct {
	Org          string
	Login        string
	AvatarUrl    string
	Followers    int64
	Following    int64
	FollowersUrl string
	FollowingUrl string
	DateCreated  time.Time
	DateModified time.Time
	DateDeleted  time.Time
}

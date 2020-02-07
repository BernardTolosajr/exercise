package models

import "time"

type Comment struct {
	MemberId     string
	MemberName   string
	Org          string
	Comment      string
	DateCreated  time.Time
	DateModified time.Time
	DateDeleted  string
}

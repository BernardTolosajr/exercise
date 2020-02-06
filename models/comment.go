package models

import "time"

type Comment struct {
	Org          string
	Comment      string
	DateCreated  time.Time
	DateModified time.Time
	DateDeleted  string
}

package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Blog Post
type Post struct {
	gorm.Model
	Title          string
	Text           string
	Author         string
	Next           *Post
	NextPostID     uint `gorm:"type:int REFERENCES posts(id) ON DELETE CASCADE"`
	Previous       *Post
	PreviousPostID uint `gorm:"type:int REFERENCES posts(id) ON DELETE CASCADE"`
	PostedOn       time.Time
	Section        []Section
	Tag            []Tag
}

// Section of Blog Post (headings)
type Section struct {
	PostID uint `gorm:"type:int REFERENCES posts(id) ON DELETE CASCADE"`
	gorm.Model
	Name string
}

// Tag of Blog Post (hashtag)
type Tag struct {
	PostID uint `gorm:"type:int REFERENCES posts(id) ON DELETE CASCADE"`
	gorm.Model
	Name string
}

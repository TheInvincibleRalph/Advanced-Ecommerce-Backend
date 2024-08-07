package models

import (
	"time"

	"gorm.io/gorm"
)

type BlogPost struct {
	gorm.Model
	ID          int       `json:"id" gorm:"primary_key,auto_increment"`
	Title       string    `json:"title" gorm:"not null"`
	Content     string    `json:"content" gorm:"not null"`
	AuthorID    int       `json:"author_id" gorm:"not null"`
	Author      User      `json:"author" gorm:"foreignKey:AuthorID"`
	PublishedAt time.Time `json:"published_at"`
	Tags        []Tag     `json:"tags" gorm:"many2many:blog_post_tags"`
}
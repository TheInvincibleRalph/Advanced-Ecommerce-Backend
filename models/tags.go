package models

import "gorm.io/gorm"

//Associates tags with blog posts.

type Tag struct {
	gorm.Model
	ID        int        `json:"id" gorm:"primary_key,auto_increment"`
	Name      string     `json:"name" gorm:"not null,index,unique"`
	BlogPosts []BlogPost `json:"blog_posts" gorm:"many2many:blog_post_tags"`
}

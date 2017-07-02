package models

import (
	"github.com/jinzhu/gorm"
)

// Article contains the Title, Author, and Body of an article
type Article struct {
	gorm.Model
	Title  string
	Author string
	Body   string
}

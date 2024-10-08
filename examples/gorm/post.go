package main

import (
	gorm2 "github.com/jace996/saas/gorm"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	gorm2.MultiTenancy
}

package repository

import (
	"todolist/domain"

	"gorm.io/gorm"
)

type UserContract interface {
	Create(conn *gorm.DB, data domain.User) (err error)
	Get(conn *gorm.DB, where map[string]any) (result domain.User, err error)
	Update(conn *gorm.DB, update map[string]any, where map[string]any) (err error)
}

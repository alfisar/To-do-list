package repository

import (
	"fmt"
	"todolist/domain"
	"todolist/internal/errorhandler"

	"gorm.io/gorm"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (obj *UserRepo) Create(conn *gorm.DB, data domain.User) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
	}

	err = conn.Debug().Table("user").Save(&data).Error
	if err != nil {
		err = fmt.Errorf("Create user error : %w", err)
	}
	return
}

func (obj *UserRepo) Get(conn *gorm.DB, where map[string]any) (result domain.User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
	}

	err = conn.Debug().Table("user").Where(where).First(&result).Error
	if err != nil {
		err = fmt.Errorf("Get user error : %w", err)
	}
	return
}

func (obj *UserRepo) Update(conn *gorm.DB, update map[string]any, where map[string]any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}
		return

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
	}
	data := conn.Debug().Table("user").Where(where).Updates(update)
	if data.Error != nil {
		err = fmt.Errorf("Update user error : %w", err)

	} else if data.RowsAffected == 0 {
		err = fmt.Errorf("Update Failed : No Rows Affected")
	}
	return
}

package repository

import "microx/srv/user/internal/domain/model"

type UserRepo interface {
	Add(user *model.User) error
	Get(id int64) (*model.User, error)
	GetByMobile(mobile string) (*model.User, error)
	Update(user *model.User) error
}

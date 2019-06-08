package repository

import "microx/srv/passport/internal/domain/model"

type UserTokenRepo interface {
	Add(user *model.UserToken) error
	GetByUserIdAndAppId(userId int64, appId int) (*model.UserToken, error)
	Update(user *model.UserToken) error
}

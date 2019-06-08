package mysql

import (
	"github.com/go-xorm/xorm"

	"microx/srv/passport/internal/domain/model"
)

type UserTokenRepo struct {
	Engine *xorm.Engine
}

func NewUserTokenRepo(engine *xorm.Engine) *UserTokenRepo {
	return &UserTokenRepo{Engine: engine}
}

func (r *UserTokenRepo) Add(v *model.UserToken) (err error) {
	if _, err = r.Engine.InsertOne(v); err != nil {
		return
	}
	return
}

func (r *UserTokenRepo) GetByUserIdAndAppId(userId int64, appId int) (v *model.UserToken, err error) {
	var (
		has bool
	)
	v = &model.UserToken{UserId: userId, AppId: appId}
	if has, err = r.Engine.Get(v); err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return v, nil
}

func (r *UserTokenRepo) Update(v *model.UserToken) (err error) {
	if _, err = r.Engine.Update(v); err != nil {
		return
	}
	return
}

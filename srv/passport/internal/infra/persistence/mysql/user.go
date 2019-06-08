package mysql

import (
	"github.com/go-xorm/xorm"

	"microx/srv/passport/internal/domain/model"
)

type UserRepo struct {
	Engine *xorm.Engine
}

func NewUserRepo(engine *xorm.Engine) *UserRepo {
	return &UserRepo{Engine: engine}
}
func (r *UserRepo) Add(user *model.User) (err error) {
	if _, err = r.Engine.InsertOne(user); err != nil {
		return
	}
	return
}

func (r *UserRepo) Get(id int64) (user *model.User, err error) {
	var (
		has bool
	)
	user = &model.User{Id: id}
	if has, err = r.Engine.Get(user); err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return user, nil
}

func (r *UserRepo) GetByMobile(mobile string) (user *model.User, err error) {
	var (
		has bool
	)
	user = &model.User{Mobile: mobile}
	if has, err = r.Engine.Get(user); err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return user, nil
}

func (r *UserRepo) Update(user *model.User) (err error) {
	if _, err = r.Engine.Update(user); err != nil {
		return
	}
	return
}

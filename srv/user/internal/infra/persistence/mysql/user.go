package mysql

import (
	"github.com/go-xorm/xorm"
	"microx/srv/user/internal/domain/model"
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

func (r *UserRepo) Get(id int64) (*model.User, error) {
	return nil, nil
}

func (r *UserRepo) GetByMobile(mobile string) (*model.User, error) {
	return nil, nil
}

func (r *UserRepo) Update(user *model.User) error {
	return nil
}

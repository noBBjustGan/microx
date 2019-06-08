package service

import (
	"microx/srv/user/internal/domain/model"
	"microx/srv/user/internal/domain/repository"
)

type UserService struct {
	userRepo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) Create(userId int64, mobile string) error {
	u := &model.User{Id: userId, Mobile: mobile}
	return s.userRepo.Add(u)
}

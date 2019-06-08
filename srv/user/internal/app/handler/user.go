package handler

import (
	"context"
	"fmt"

	user "microx/srv/user/api"
	"microx/srv/user/internal/app/service"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Create(ctx context.Context, req *user.Request, rsp *user.Response) (err error) {
	if err = h.svc.Create(req.UserId, req.Mobile); err != nil {
		fmt.Println("xxxxxxxxxxxx")
		return
	}
	fmt.Println(req.UserId, req.Mobile)
	rsp = &user.Response{}
	return
}

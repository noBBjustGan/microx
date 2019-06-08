package handler

import (
	"context"
	"microx/pkg/log"

	"microx/common"
	"microx/common/typ"
	passport "microx/srv/passport/api"
	"microx/srv/passport/internal/app/service"
)

type PassportHandler struct {
	svc *service.PassportService
}

func NewPassportHandler(svc *service.PassportService) *PassportHandler {
	return &PassportHandler{svc: svc}
}

func (h *PassportHandler) Sms(ctx context.Context, req *passport.Request, rsp *passport.Response) error {
	log.Info("Sms: mobile=%s", req.Mobile)
	// TODO: 通过短信服务获取验证码
	rsp.Code = "8888"
	return nil
}

func (h *PassportHandler) SmsLogin(ctx context.Context, req *passport.SmsLoginRequest, rsp *passport.SmsLoginResponse) (err error) {
	log.Info("SmsLogin: mobile=%s code=%s", req.Mobile, req.Code)
	var (
		header *typ.Header
	)
	if header, err = common.GetHeaderFromContext(ctx); err != nil {
		return
	}
	rsp.TokenInfo, err = h.svc.SmsLogin(ctx, req.Mobile, req.Code, header.AppId)
	return
}

func (h *PassportHandler) Login(ctx context.Context, req *passport.LoginRequest, rsp *passport.LoginResponse) (err error) {
	log.Info("Login: mobile=%s passwd=%s", req.Mobile, req.Passwd)
	var (
		header *typ.Header
	)
	header, err = common.GetHeaderFromContext(ctx)
	if err != nil {
		return err
	}
	if rsp.TokenInfo, err = h.svc.Login(ctx, req.Mobile, req.Passwd, header.AppId); err != nil {
		return err
	}
	return
}

func (h *PassportHandler) OAuthLogin(ctx context.Context, req *passport.OAuthLoginRequest, rsp *passport.OAuthLoginResponse) error {
	return nil
}

func (h *PassportHandler) ValidateToken(ctx context.Context, req *passport.TokenRequest, rsp *passport.TokenResponse) error {
	log.Info("ValidateToken")
	header, err := common.GetHeaderFromContext(ctx)
	if err != nil {
		return err
	}
	log.Infof("ValidateToken: uid=%d token=%s", header.UserId, header.Token)
	return h.svc.ValidateToken(ctx, header.UserId, header.Token)
}

func (h *PassportHandler) SetPwd(ctx context.Context, req *passport.SetPwdRequest, rsp *passport.SetPwdResponse) error {
	log.Info("SetPwd")
	header, err := common.GetHeaderFromContext(ctx)
	if err != nil {
		return err
	}

	if rsp.TokenInfo, err = h.svc.SetPwd(ctx, header.UserId, req.Passwd, header.AppId); err != nil {
		return err
	}
	return nil
}

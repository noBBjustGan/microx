package controller

import (
	"microx/api/xdd/factory"
	"microx/pkg/context"
	"microx/pkg/log"
	passport "microx/srv/passport/api"
)

// PassportController ...
type PassportController struct {
}

func (this *PassportController) Login(mctx *context.MxContext) {
	log.Info("Login...")
	request := &passport.LoginRequest{}
	if err := mctx.ParseJSON(request); err != nil {
		log.Error(err)
		mctx.ResponseError(err)
		return
	}
	// TODO: request相关参数校验
	response, err := factory.PassportClient.Login(toContext(mctx), request)
	if err != nil {
		log.Error(err)
		mctx.ResponseError(err)
		return
	}

	mctx.Response(response)
	return
}

func (this *PassportController) Sms(mctx *context.MxContext) {
	log.Info("Sms...")
	request := &passport.Request{}
	if err := mctx.ParseJSON(request); err != nil {
		log.Error(err)
		mctx.ResponseError(err)
		return
	}
	// TODO: request相关参数校验
	response, err := factory.PassportClient.Sms(toContext(mctx), request)
	if err != nil {
		log.Error(err)
		mctx.ResponseError(err)
		return
	}

	mctx.Response(response)
	return
}

func (this *PassportController) SmsLogin(mctx *context.MxContext) {
	log.Info("SmsLogin...")
	request := &passport.SmsLoginRequest{}
	if err := mctx.ParseJSON(request); err != nil {
		log.Error(err)
		mctx.ResponseError(err)
		return
	}

	// TODO: request相关参数校验

	response, err := factory.PassportClient.SmsLogin(toContext(mctx), request)
	if err != nil {
		log.Error(err)
		mctx.ResponseError(err)
		return
	}

	mctx.Response(response)
	return
}

func (this *PassportController) OauthLogin(mctx *context.MxContext) {
	log.Info("OauthLogin...")
	request := &passport.OAuthLoginRequest{}
	if err := mctx.ParseJSON(request); err != nil {
		log.Error(err)
		mctx.ResponseError(err)
		return
	}
	// TODO: request相关参数校验
	response, err := factory.PassportClient.OAuthLogin(toContext(mctx), request)
	if err != nil {
		log.Error(err)
		mctx.ResponseError(err)
		return
	}

	mctx.Response(response)

	return
}

func (this *PassportController) SetPwd(mctx *context.MxContext) {
	log.Info("SetPwd...")
	request := &passport.SetPwdRequest{}
	if err := mctx.ParseJSON(request); err != nil {
		log.Error(err)
		mctx.ResponseError(err)
		return
	}
	// TODO: request相关参数校验
	response, err := factory.PassportClient.SetPwd(toContext(mctx), request)
	if err != nil {
		log.Error(err)
		mctx.ResponseError(err)
		return
	}

	mctx.Response(response)

	return
}

package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"microx/common/constant"
	"microx/common/errors"
	topic "microx/common/proto/topic"
	"microx/pkg/capx"
	"microx/pkg/log"
	gid "microx/srv/gid/api"
	passport "microx/srv/passport/api"
	"microx/srv/passport/internal/domain/model"
	"microx/srv/passport/internal/domain/repository"
	user "microx/srv/user/api"
)

type PassportService struct {
	engine        *xorm.Engine
	userClient    user.UserService
	gidClient     gid.GidService
	userRepo      repository.UserRepo
	userTokenRepo repository.UserTokenRepo
}

func NewPassportService(
	engine *xorm.Engine,
	gidClient gid.GidService,
	userClient user.UserService,
	repo repository.UserRepo,
	userTokenRepo repository.UserTokenRepo) *PassportService {
	return &PassportService{
		engine:        engine,
		gidClient:     gidClient,
		userClient:    userClient,
		userRepo:      repo,
		userTokenRepo: userTokenRepo,
	}
}

func (p *PassportService) SmsLogin(ctx context.Context, mobile, code string, appId int) (token *passport.TokenInfo, err error) {
	var (
		user *model.User
	)
	// TODO: 校验验证码
	if user, err = p.userRepo.GetByMobile(mobile); err != nil {
		log.Error(err)
		return
	}

	if user == nil {
		token, err = p.register(ctx, mobile)
	} else {
		if token, err = p.updateToken(ctx, user.Id, appId); err != nil {
			return nil, err
		}
	}

	return
}

func (p *PassportService) Login(ctx context.Context, mobile, passwd string, appId int) (token *passport.TokenInfo, err error) {
	var (
		user *model.User
	)
	if user, err = p.userRepo.GetByMobile(mobile); err != nil {
		log.Error(err)
		return nil, err
	}
	if user == nil {
		return nil, errors.New(errors.ErrUserNotExists, "")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(passwd)); err != nil {
		return nil, errors.New(errors.ErrPasswordError, "")
	}
	if token, err = p.updateToken(ctx, user.Id, appId); err != nil {
		return nil, err
	}
	return
}

func (p *PassportService) ValidateToken(ctx context.Context, userId int64, token string) error {
	return nil
}

func (p *PassportService) SetPwd(ctx context.Context, userId int64, passwd string, appId int) (token *passport.TokenInfo, err error) {
	var (
		user       *model.User
		userToken  *model.UserToken
		passwdHash []byte
	)
	if user, err = p.userRepo.Get(userId); err != nil {
		return
	}

	if user == nil {
		return nil, errors.New(errors.ErrUserNotExists, "")
	}

	if passwdHash, err = bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost); err != nil {
		return
	}
	user.Passwd = string(passwdHash)

	if userToken, err = p.userTokenRepo.GetByUserIdAndAppId(userId, appId); err != nil {
		return
	}
	userToken.AccessToken = uuid.New().String()
	userToken.RefreshToken = uuid.New().String()

	session := p.engine.NewSession()
	defer func() {
		session.Close()
	}()

	if err = session.Begin(); err != nil {
		log.Error(err)
		return
	}

	if _, err = session.Table(new(model.User)).ID(user.Id).Update(user); err != nil {
		session.Rollback()
		log.Error(err)
		return
	}

	if _, err = session.Table(new(model.UserToken)).ID(userToken.Id).Update(userToken); err != nil {
		session.Rollback()
		log.Error(err)
		return
	}

	if err = session.Commit(); err != nil {
		log.Error(err)
		return
	}

	token = &passport.TokenInfo{
		UserId:       userToken.UserId,
		Token:        userToken.AccessToken,
		RefreshToken: userToken.RefreshToken,
		ExpiredAt:    time.Now().Unix() + 8640000,
	}

	return
}

func (p *PassportService) updateToken(ctx context.Context, userId int64, appId int) (token *passport.TokenInfo, err error) {
	var (
		userToken *model.UserToken
	)
	userToken, err = p.userTokenRepo.GetByUserIdAndAppId(userId, appId)
	if err != nil {
		return
	}

	userToken.AccessToken = uuid.New().String()
	userToken.RefreshToken = uuid.New().String()
	userToken.ExpiresIn = 8640000
	p.userTokenRepo.Update(userToken)

	token = &passport.TokenInfo{
		UserId:       userToken.UserId,
		Token:        userToken.AccessToken,
		RefreshToken: userToken.RefreshToken,
		ExpiredAt:    time.Now().Unix() + 8640000,
	}

	return
}

func (p *PassportService) register(ctx context.Context, mobile string) (*passport.TokenInfo, error) {
	rsp, err := p.gidClient.GetMulti(ctx, &gid.MultiRequest{Count: 3})
	if err != nil {
		return nil, err
	}

	if len(rsp.Ids) != 3 {
		return nil, fmt.Errorf("the number of ids dose not match expected %v got %v", 3, len(rsp.Ids))
	}

	session := p.engine.NewSession()
	defer func() {
		session.Close()
	}()

	if err := session.Begin(); err != nil {
		log.Error(err)
		return nil, err
	}

	// user表
	u := &model.User{
		Id:     rsp.Ids[0],
		Mobile: mobile,
	}
	if _, err := session.InsertOne(u); err != nil {
		session.Rollback()
		log.Error(err)
		return nil, err
	}

	u1 := uuid.New().String()
	u2 := uuid.New().String()
	ut := &model.UserToken{
		Id:           rsp.Ids[1],
		AppId:        1,
		UserId:       u.Id,
		ExpiresIn:    8640000,
		AccessToken:  u1,
		RefreshToken: u2,
	}

	if _, err := session.InsertOne(ut); err != nil {
		session.Rollback()
		log.Error(err)
		return nil, err
	}

	// 分布式事务处理
	// 发布topic.user.created事件
	msg := &topic.UserCreated{
		Id:    rsp.Ids[2],
		Topic: constant.TOPIC_USER_CREATED,
		Info:  &topic.UserInfo{UserId: u.Id, Mobile: mobile},
	}

	if err = capx.StorePublished(session, rsp.Ids[2], constant.TOPIC_USER_CREATED, msg); err != nil {
		session.Rollback()
		log.Error(err)
	}

	if err := session.Commit(); err != nil {
		log.Error(err)
		return nil, err
	}

	capx.Publish(rsp.Ids[2], constant.TOPIC_USER_CREATED, msg)

	tokenInfo := &passport.TokenInfo{
		UserId:       u.Id,
		Token:        u1,
		RefreshToken: u2,
		ExpiredAt:    ut.UpdatedAt.Unix() + ut.ExpiresIn,
	}
	return tokenInfo, nil
}

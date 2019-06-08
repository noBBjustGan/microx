package handler

import (
	"context"
	"github.com/go-xorm/xorm"
	"github.com/golang/protobuf/proto"
	topic "microx/common/proto/topic"
	"microx/pkg/capx"
	"microx/pkg/log"
	"microx/srv/user/internal/app/service"
)

type Subscriber struct {
	engine *xorm.Engine
	svc    *service.UserService
}

func NewSubscriber(engine *xorm.Engine, svc *service.UserService) *Subscriber {
	return &Subscriber{engine: engine, svc: svc}
}

func (sub *Subscriber) UserCreated() func(ctx context.Context, msg *topic.UserCreated) (err error) {
	return func(ctx context.Context, msg *topic.UserCreated) (err error) {
		log.Infof("received msg topic=%s", msg.Topic)
		if err = capx.StoreReceived(msg.Id, msg.Topic, msg); err != nil {
			// 重复的消息，保存失败，从而保证幂等性
			return
		}

		if err = sub.svc.Create(msg.Info.UserId, msg.Info.Mobile); err != nil {
			capx.Consumed(msg.Id, 2)
			return
		}

		capx.Consumed(msg.Id, 1)
		return
	}

}

func (sub *Subscriber) CapxUserCreated() capx.ConsumerFn {
	return func(pb proto.Message) (err error) {
		log.Info("CapxUserCreated")
		msg := pb.(*topic.UserCreated)
		return sub.svc.Create(msg.Info.UserId, msg.Info.Mobile)
	}
}

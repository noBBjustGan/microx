package capx

import (
	"context"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"microx/pkg/capx/model"
	"microx/pkg/log"
)

var engine *xorm.Engine

func StorePublished(session *xorm.Session, id int64, topic string, pb interface{}) (err error) {
	var msg []byte
	if msg, err = proto.Marshal(pb.(proto.Message)); err != nil {
		return err
	}

	name := proto.MessageName(pb.(proto.Message))
	v := model.Published{
		Id:     id,
		Topic:  topic,
		Name:   name,
		Msg:    msg,
		Status: 0,
	}

	if _, err = session.InsertOne(&v); err != nil {
		return
	}

	return
}

func StoreReceived(id int64, topic string, pb interface{}) (err error) {
	if ok, _ := engine.Exist(&model.Received{Id: id}); ok {
		return fmt.Errorf("msg id=%d exist", id)
	}
	var msg []byte
	if msg, err = proto.Marshal(pb.(proto.Message)); err != nil {
		return err
	}

	name := proto.MessageName(pb.(proto.Message))
	v := model.Received{
		Id:     id,
		Topic:  topic,
		Name:   name,
		Msg:    msg,
		Status: 0,
	}

	if _, err = engine.InsertOne(&v); err != nil {
		return
	}

	return
}

func Publish(id int64, topic string, msg interface{}) error {
	log.Infof("publish topic %s", topic)

	p := micro.NewPublisher(topic, client.DefaultClient)
	if err := p.Publish(context.Background(), msg); err != nil {
		log.Error("publish err:", err)
		updatePublished(id, map[string]interface{}{"status": 2})
		return err
	} else {
		log.Infof("Published %v\n", msg)
		updatePublished(id, map[string]interface{}{"status": 1})
		return nil
	}
}

func Init(e *xorm.Engine) {
	engine = e
	go sending()
	go consuming()
}

func updatePublished(id int64, m map[string]interface{}) error {
	_, err := engine.Table(new(model.Published)).ID(id).Update(m)
	return err
}

func updateReceived(id int64, m map[string]interface{}) error {
	_, err := engine.Table(new(model.Received)).ID(id).Update(m)
	return err
}

func Consumed(id int64, status int) {
	updateReceived(id, map[string]interface{}{"status": status})
}

type ConsumerFn func(proto.Message) error

var consumers = map[string]ConsumerFn{}

func RegisterConsumer(name string, fn ConsumerFn) {
	consumers[name] = fn
}

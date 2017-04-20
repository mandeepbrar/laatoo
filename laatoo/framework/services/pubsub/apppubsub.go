package pubsub

import (
	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//"laatoo/sdk/errors"
)

type AppPubSubService struct {
	subscribers map[string][]core.ServiceFunc
}

func (svc *AppPubSubService) Publish(ctx core.RequestContext, topic string, message interface{}) error {
	subs, ok := svc.subscribers[topic]
	if ok {
		for _, sub := range subs {
			go func(ctx core.RequestContext, lstnr core.ServiceFunc, data interface{}) {
				req := ctx.SubContext("Message")
				req.Set("messagetype", topic)
				req.SetRequest(data)
				lstnr(req)
			}(ctx, sub, message)
		}
	}
	return nil
}

func (svc *AppPubSubService) Subscribe(ctx core.ServerContext, topics []string, lstnr core.ServiceFunc) error {
	for _, topic := range topics {
		subs, prs := svc.subscribers[topic]
		if !prs {
			subs = []core.ServiceFunc{}
		}
		subs = append(subs, lstnr)
		svc.subscribers[topic] = subs
	}
	return nil
}

func (svc *AppPubSubService) Initialize(ctx core.ServerContext, conf config.Config) error {
	svc.subscribers = make(map[string][]core.ServiceFunc, 10)
	return nil
}

func (svc *AppPubSubService) Invoke(ctx core.RequestContext) error {
	return nil
}
func (svc *AppPubSubService) Start(ctx core.ServerContext) error {
	return nil
}

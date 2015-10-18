package service

import (
	"github.com/labstack/echo"
)

type TopicListener func(ctx *echo.Context, topic string, message interface{})

type PubSub interface {
	Publish(ctx *echo.Context, topic string, message interface{}) error
	Subscribe(ctx *echo.Context, topics []string, lstnr TopicListener) error
}

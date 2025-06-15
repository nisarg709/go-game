package services

import (
	"github.com/globalsign/mgo"
	"errors"
	"encoding/json"
)

type Handler func(context *Context, event *Event) error

type Context struct {
	Mongo *mgo.Session
}

type Event struct {
	Action string `json:"event"`
	Data json.RawMessage `json:"data"`
}

type Service struct {
	Context  *Context
	handlers map[string]Handler
}

func (self *Service) AddHandler(action string, handler Handler) {

	if self.handlers == nil {
		self.handlers = make(map[string]Handler)
	}

	self.handlers[action] = handler
}

func (self *Service) Work(event *Event) (error) {

	handler, ok := self.handlers[event.Action];
	if !ok {
		return errors.New("no handler for event with action " + event.Action)
	}

	err := handler(self.Context, event)

	if err != nil {
		return err
	}

	return nil;
}

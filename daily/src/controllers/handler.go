package controllers

import (
	"eman/passport/daily/src/api"
	"eman/passport/daily/src/repositories"
	"github.com/globalsign/mgo"
)

type Handler struct {
	Container *Container
	caller    *api.Caller
}

type Container struct {
	LaunchesRepository *repositories.LaunchesRepository
	RewardsRepository  *repositories.RewardsRepository
	EventsRepository   *repositories.EventsRepository
}

func New(mongo *mgo.Session) *Handler {

	parent := &repositories.MongoRepository{mongo}

	// Handler
	return &Handler{
		Container: &Container{
			&repositories.LaunchesRepository{parent},
			&repositories.RewardsRepository{parent},
			&repositories.EventsRepository{parent},
		},
	}
}

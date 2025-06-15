package bootstrap

import (
	"github.com/globalsign/mgo"
	"eman/passport/achievements/src/handlers"
	"eman/passport/achievements/src/services"
)

func NewService() services.Service {

	// DB Connect
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err);
	}

	// Build Service
	service := services.Service{
		Context: &services.Context{
			Mongo: session,
		},
	}

	service.AddHandler("survey.completed", handlers.SurveyCompletedHandler);

	return service
}

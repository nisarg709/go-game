package service

import "eman/passport/sync/src/app"

func Bootstrap() *app.Server {

	configuration := NewConfiguration()

	config := app.ServerConfig{
		LogExecutionTime: configuration.LogExecution,
		Port: configuration.Port,
	}

	service:= NewService(configuration)

	server := app.Server{
		config,
		service,
		nil,
	}

	server.Load()

	return &server
}

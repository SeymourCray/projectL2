package app

import (
	"dev11/config"
	"dev11/internal/controller"
	"dev11/internal/repository"
	"dev11/internal/usecase"
	"fmt"
	"log"
	"net/http"
)

func Run(config *config.Config) {
	eventsRepository := repository.NewCashedEvents()
	eventUseCase := usecase.NewEventUseCase(eventsRepository)
	eventController := controller.NewEventController(eventUseCase)
	eventController.RegisterHandlers()

	log.Printf("Starting HTTP server on port %s\n", config.Port)

	port := fmt.Sprintf(":%s", config.Port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln(err)
	}
}

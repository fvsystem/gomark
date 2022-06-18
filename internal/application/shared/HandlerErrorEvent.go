package shared

import (
	"log"
)

type HandlerError struct {
}

func (e *HandlerError) Run(event Event) {
	errorText, ok := event.Data.(string)
	if !ok {
		log.Fatal("HandlerError: event.Data is not string")
	}
	log.Fatalf("Error %s, ", errorText)
}

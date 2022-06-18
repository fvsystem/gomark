package shared

import (
	"log"
)

type HandlerError struct {
}

func (e *HandlerError) Run(event Event) {
	errorText, ok := event.Data.(string)
	if !ok {
		panic("HandlerError: event.Data is not string")
	}
	log.Panicf("Error %s, ", errorText)
}

package shared

func NewEventError(errorMessage string) Event {
	return Event{
		Name: "Error",
		Data: errorMessage,
	}
}

package event

import "github.com/fvsystem/gomark/internal/application/shared"

var StopTestEvent shared.Event = shared.Event{
	Name: "StopTest",
	Data: interface{}(nil),
}

package event

import "github.com/fvsystem/gomark/internal/application/shared"

var StartTestEvent shared.Event = shared.Event{
	Name: "StartTest",
	Data: interface{}(nil),
}

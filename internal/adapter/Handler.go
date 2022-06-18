package adapter

import "github.com/fvsystem/gomark/internal/application/shared"

type Handler interface {
	Run(event shared.Event)
}

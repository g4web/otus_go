package messagequeue

import (
	"context"
)

type HandleFunction func(context.Context, <-chan *EventDTO)

type MessageQueue interface {
	Open() error
	Close()
	Publish(event *EventDTO) error
	Consume(ctx context.Context, handleFunc HandleFunction) error
}

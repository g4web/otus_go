package messagequeue

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	addr          string
	queueName     string
	handlersCount int
	conn          *amqp.Connection
	channel       *amqp.Channel
	logger        logger.Logger
}

func NewRabbitMQ(addr string, queueName string, handlersCount int, logger logger.Logger) *RabbitMQ {
	return &RabbitMQ{addr: addr, queueName: queueName, handlersCount: handlersCount, logger: logger}
}

func (r *RabbitMQ) Open() error {
	var err error
	r.conn, err = amqp.Dial(r.addr)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	r.channel, err = r.conn.Channel()
	if err != nil {
		r.logger.Error("fail open channel")
	}

	_, err = r.channel.QueueDeclare(r.queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	r.logger.Info("connected to rabbitmq")

	return nil
}

func (r *RabbitMQ) Close() {
	if err := r.channel.Cancel("", true); err != nil {
		r.logger.Error(err.Error())
	}
	if err := r.conn.Close(); err != nil {
		r.logger.Error(err.Error())
	}
}

func (r *RabbitMQ) Publish(event *EventDTO) error {
	encodedData, err := json.Marshal(event)
	if err != nil {
		r.logger.Error("fail encode event")
	}

	return r.channel.PublishWithContext(
		context.Background(),
		"",
		r.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        encodedData,
		},
	)
}

func (r *RabbitMQ) Consume(ctx context.Context, handleFunc HandleFunction) error {
	for {
		rabbitMqChannel, err := r.announceQueue()
		if err != nil {
			return err
		}

		eventChannel := make(chan *EventDTO)

		go func(rabbitMqChannel <-chan amqp.Delivery) {
			for {
				msg, ok := <-rabbitMqChannel
				if !ok {
					break
				}
				event := &EventDTO{}
				if err := json.Unmarshal(msg.Body, event); err != nil {
					r.logger.Error("event unmarshal failed")
					break
				}

				if err := msg.Ack(false); err != nil {
					r.logger.Error("ack failed")
					break
				}
				eventChannel <- event
			}
		}(rabbitMqChannel)

		for i := 0; i < r.handlersCount; i++ {
			go handleFunc(ctx, eventChannel)
		}
	}
}

func (r *RabbitMQ) announceQueue() (<-chan amqp.Delivery, error) {
	msgs, err := r.channel.Consume(
		r.queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return msgs, nil
}

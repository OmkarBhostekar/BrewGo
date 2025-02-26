package rabbitmq

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	exchange string
}

func NewRabbitMQ(exchange, address string) (*RabbitMQ, error) {
	address = "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(address)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	err = ch.ExchangeDeclare(
		exchange, "direct", true, false, false, false, nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{conn, ch, exchange}, nil
}

func (r *RabbitMQ) Publish(routingKey string, message string) error {
	return r.channel.Publish(
		r.exchange, routingKey, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(message),
			Timestamp:   time.Now(),
		},
	)
}

func (r *RabbitMQ) Consume(queueName, routingKey string, handler func(msg amqp.Delivery)) error {
	q, err := r.channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = r.channel.QueueBind(q.Name, routingKey, r.exchange, false, nil)
	if err != nil {
		return err
	}

	// Qos is used to limit the number of unacknowledged messages on a channel
	r.channel.Qos(1, 0, false)
	messages, err := r.channel.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for msg := range messages {
			handler(msg)
		}
	}()

	log.Printf("Listening for %s on queue %s", routingKey, queueName)
	return nil
}

func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.conn.Close()
}

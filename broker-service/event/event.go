package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	durable := true
	autoDelete := false
	internal := false
	noWait := false
	args := amqp.Table{}

	return ch.ExchangeDeclare(
		"logs_topic",
		"topic",
		durable,
		autoDelete,
		internal,
		noWait,
		args,
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	durable := false
	autoDelete := true
	internal := false
	noWait := false
	args := amqp.Table{}

	return ch.QueueDeclare(
		"",
		durable,
		autoDelete,
		internal,
		noWait,
		args,
	)
}

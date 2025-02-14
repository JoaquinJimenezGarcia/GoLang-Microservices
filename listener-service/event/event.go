package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // Name of the exchange
		"topic",      // Type of exchange
		true,         // Durable
		false,        // Not autodelete
		false,        // not only internal
		false,        // no waiting
		nil,          // Arguments
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // Name
		false, // Not durable
		false, // Not autodelete
		true,  // Exclusive
		false, // No wait
		nil,   // Arguments
	)
}

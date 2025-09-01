package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}

	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	return declareExchange(channel)
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (consumer *Consumer) Listen(topics []string) error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	queue, err := declareRandomQueue(channel)
	if err != nil {
		return err
	}

	noWait := false
	args := amqp.Table{}
	for _, s := range topics {
		err = channel.QueueBind(
			queue.Name,
			s,
			"logs_topic",
			noWait,
			args,
		)
		if err != nil {
			return err
		}
	}

	autoAck := true
	exclusive := false
	noLocal := false
	consumeNoWait := false
	consumeArgs := amqp.Table{}
	messages, err := channel.Consume(
		queue.Name,
		"",
		autoAck,
		exclusive,
		noLocal,
		consumeNoWait,
		consumeArgs,
	)

	forever := make(chan bool)
	go func() {
		for d := range messages {
			var payload Payload
			_ = json.Unmarshal(d.Body, &payload)

			go handlePayload(payload)
		}
	}()

	fmt.Printf("Waiting for message (Exchange, Queue) [logs_topic, %s]\n", queue.Name)
	<-forever

	return nil
}

func handlePayload(payload Payload) {
	switch payload.Name {
	case "log", "event":
		err := logEvent(payload)
		fmt.Printf("Log: %s\n", payload.Data)
		if err != nil {
			log.Println(err)
			return
		}
	case "auth":
		fmt.Printf("Auth: %s\n", payload.Data)
	default:
		err := logEvent(payload)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func logEvent(event Payload) error {
	jsonData, _ := json.MarshalIndent(event, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://logger-service/log", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return err
	}

	return nil
}

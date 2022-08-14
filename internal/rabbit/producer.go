package rabbit

import (
	"bannerRotator/internal/config"
	"bannerRotator/internal/logger"
	"github.com/streadway/amqp"
)

type Producer struct {
	logger     *logger.Logger
	config     config.RabbitConfig
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewProducer(
	logger *logger.Logger,
	config config.RabbitConfig,
) *Producer {
	return &Producer{
		logger: logger,
		config: config,
	}
}

func (p *Producer) Connect() error {
	connection, err := amqp.Dial(p.config.CreateDSN())
	if err != nil {
		return err
	}

	p.connection = connection

	channel, err := p.connection.Channel()
	if err != nil {
		return err
	}

	p.channel = channel

	_, err = p.channel.QueueDeclare(
		p.config.Queue,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *Producer) Disconnect() error {
	if p.connection == nil {
		return nil
	}

	return p.connection.Close()
}

func (p *Producer) Publish(body []byte) error {
	if p.connection == nil || p.connection.IsClosed() {
		if err := p.Connect(); err != nil {
			return err
		}
	}

	return p.channel.Publish("", p.config.Queue, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
}

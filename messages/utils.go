package messages

import (
	"github.com/mailru/easyjson"
	"github.com/streadway/amqp"

	"github.com/studtool/common/consts"
)

func (c *MqClient) declareQueue(queueName string) (amqp.Queue, error) {
	return c.channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
}

func (c *MqClient) runConsumer(queueName string, handler messageHandler) error {
	messages, err := c.channel.Consume(
		queueName,
		consts.EmptyString,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range messages {
			handler(d.Body)
		}
	}()

	return nil
}

func (c *MqClient) unmarshalMessageBody(data []byte, v easyjson.Unmarshaler) error {
	return easyjson.Unmarshal(data, v)
}

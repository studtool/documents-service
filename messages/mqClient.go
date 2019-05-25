package messages

import (
	"fmt"

	"github.com/streadway/amqp"
	"go.uber.org/dig"

	"github.com/studtool/common/logs"
	"github.com/studtool/common/queues"
	"github.com/studtool/common/utils"

	"github.com/studtool/documents-service/config"
	"github.com/studtool/documents-service/logic"
	"github.com/studtool/documents-service/utils"
)

type MqClient struct {
	connStr    string
	connection *amqp.Connection

	channel *amqp.Channel

	documentUsersToCreateQueue amqp.Queue
	documentUsersToDeleteQueue amqp.Queue

	usersService logic.UsersService

	structLogger  logs.Logger
	reflectLogger logs.Logger
}

type MqClientParams struct {
	dig.In

	UsersService logic.UsersService
}

func NewMqClient(params MqClientParams) *MqClient {
	c := &MqClient{
		connStr: fmt.Sprintf("amqp://%s:%s@%s:%d/",
			config.MqUser.Value(), config.MqPassword.Value(),
			config.MqHost.Value(), config.MqPort.Value(),
		),
		usersService: params.UsersService,
	}

	c.structLogger = srvutils.MakeStructLogger(c)
	c.reflectLogger = srvutils.MakeReflectLogger(c)

	c.structLogger.Info("initialized")

	return c
}

func (c *MqClient) OpenConnection() error {
	var conn *amqp.Connection
	err := utils.WithRetry(func(n int) (err error) {
		if n > 0 {
			c.structLogger.Infof("opening message queue connection. retry #%d", n)
		}
		conn, err = amqp.Dial(c.connStr)
		return err
	}, config.MqConnNumRet.Value(), config.MqConnRetItv.Value())
	if err != nil {
		return err
	}

	c.connection = conn

	c.channel, err = conn.Channel()
	if err != nil {
		return err
	}

	c.documentUsersToCreateQueue, err =
		c.declareQueue(queues.DocumentUsersToCreateQueueName)
	if err != nil {
		return err
	}

	c.documentUsersToDeleteQueue, err =
		c.declareQueue(queues.DocumentUsersToDeleteQueueName)
	if err != nil {
		return err
	}

	c.structLogger.Info("connection opened")

	return nil
}

func (c *MqClient) CloseConnection() error {
	if err := c.channel.Close(); err != nil {
		return err
	}
	if err := c.connection.Close(); err != nil {
		return err
	}

	c.structLogger.Info("connection closed")

	return nil
}

type messageHandler func(data []byte)

func (c *MqClient) Run() error {
	err := c.runConsumer(
		queues.DocumentUsersToCreateQueueName,
		c.createDocumentUser,
	)
	if err != nil {
		return err
	}

	err = c.runConsumer(
		queues.DocumentUsersToDeleteQueueName,
		c.deleteDocumentUser,
	)
	if err != nil {
		return err
	}

	c.structLogger.Infof("ready to consume messages")

	return nil
}

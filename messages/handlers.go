package messages

import (
	"github.com/studtool/common/queues"
	"github.com/studtool/common/types"

	"github.com/studtool/documents-service/models"
)

func (c *MqClient) createDocumentUser(body []byte) {
	data := &queues.DocumentUserToCreateData{}
	if err := c.unmarshalMessageBody(body, data); err != nil {
		c.structLogger.Errorf(
			"wrong message received [queue = '%s]",
			queues.DocumentUsersToCreateQueueName,
		)
	} else {
		u := &models.User{
			ID: types.ID(data.UserID),
		}
		if err := c.usersService.AddUser(u); err == nil {
			c.structLogger.Infof("user [id = '%s'] created", u.ID)
		} else {
			c.structLogger.Errorf("user [id = '%s'] not created", u.ID)
		}
	}
}

func (c *MqClient) deleteDocumentUser(body []byte) {
	data := &queues.DocumentUserToDeleteData{}
	if err := c.unmarshalMessageBody(body, data); err != nil {
		c.structLogger.Errorf(
			"wrong message received [queue = '%s]",
			queues.DocumentUsersToDeleteQueueName,
		)
	} else {
		u := &models.User{
			ID: types.ID(data.UserID),
		}
		if err := c.usersService.DeleteUser(u); err == nil {
			c.structLogger.Infof("user [id = '%s'] deleted", u.ID)
		} else {
			c.structLogger.Errorf("user [id = '%s'] not deleted", u.ID)
		}
	}
}

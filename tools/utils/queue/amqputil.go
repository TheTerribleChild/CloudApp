package queueutil

import (
	"github.com/streadway/amqp"
	"fmt"
	"errors"
)

type AmqpQueueBuilder struct{
	host string
	user string
	password string
	port int
	queueName string
	durable bool
	autoDelete bool
	exclusive bool
	noWait bool
	args map[string]interface{}
	passiveDeclare bool
}

func GetAMQPQueueBuilder(host string, user string, password string, queueName string) *AmqpQueueBuilder {
	return &AmqpQueueBuilder{host:host, user:user, password:password, queueName: queueName}
}

func(instance *AmqpQueueBuilder) SetPort(port int) *AmqpQueueBuilder{
	instance.port = port
	return instance
}

func(instance *AmqpQueueBuilder) SetDurable(durable bool) *AmqpQueueBuilder{
	instance.durable = durable
	return instance
}

func(instance *AmqpQueueBuilder) SetAutoDelete(autoDelete bool) *AmqpQueueBuilder{
	instance.autoDelete = autoDelete
	return instance
}

func(instance *AmqpQueueBuilder) SetNoWait(noWait bool) *AmqpQueueBuilder{
	instance.noWait = noWait
	return instance
}

func(instance *AmqpQueueBuilder) SetDeclarePassive(passiveDeclare bool) *AmqpQueueBuilder{
	instance.passiveDeclare = passiveDeclare
	return instance
}

func(instance *AmqpQueueBuilder) SetArguements(args map[string]interface{}) *AmqpQueueBuilder{
	instance.args = args
	return instance
}

func(instance *AmqpQueueBuilder) Build() (*amqp.Connection, *amqp.Channel, amqp.Queue, error){
	if len(instance.host) == 0 || len(instance.user) == 0 || len(instance.password) == 0 || len(instance.queueName) == 0 {
		return nil, nil, amqp.Queue{}, errors.New("missing required arguements")
	}
	if instance.port == 0 {
		instance.port = 5672
	} else if instance.port < 0 || instance.port > 65535 {
		return nil, nil, amqp.Queue{}, errors.New("invalid amqp port number")
	}
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%d/", instance.user, instance.password, instance.host, instance.port)
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, nil, amqp.Queue{}, err
	}
	channel, err := connection.Channel()
	if err != nil {
		return nil, nil, amqp.Queue{}, err
	}
	var queue amqp.Queue
	if instance.passiveDeclare {
		queue , err = channel.QueueDeclarePassive(
			instance.queueName, 
			instance.durable, 
			instance.autoDelete, 
			instance.exclusive, 
			instance.noWait, 
			instance.args,
		)
	}else{
		queue, err = channel.QueueDeclare(
			instance.queueName, 
			instance.durable, 
			instance.autoDelete, 
			instance.exclusive, 
			instance.noWait, 
			instance.args,
		)
	}
	return connection, channel, queue, err
}
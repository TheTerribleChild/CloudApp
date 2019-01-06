package queueutil

import (
	"github.com/streadway/amqp"
	"fmt"
	"errors"
)

type AmqpQueueBuilder struct{
	Host string
	User string
	Password string
	Port int
	QueueName string
	Durable bool
	AutoDelete bool
	Exclusive bool
	NoWait bool
	Args map[string]interface{}
	PassiveDeclare bool
}

func(instance *AmqpQueueBuilder) Build() (*amqp.Connection, *amqp.Channel, amqp.Queue, error){
	if len(instance.Host) == 0 || len(instance.User) == 0 || len(instance.Password) == 0 || len(instance.QueueName) == 0 {
		return nil, nil, amqp.Queue{}, errors.New("missing required arguements")
	}
	if instance.Port == 0 {
		instance.Port = 5672
	} else if instance.Port < 0 || instance.Port > 65535 {
		return nil, nil, amqp.Queue{}, errors.New("invalid amqp Port number")
	}
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%d/", instance.User, instance.Password, instance.Host, instance.Port)
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, nil, amqp.Queue{}, err
	}
	channel, err := connection.Channel()
	if err != nil {
		return nil, nil, amqp.Queue{}, err
	}
	var queue amqp.Queue
	if instance.PassiveDeclare {
		queue , err = channel.QueueDeclarePassive(
			instance.QueueName, 
			instance.Durable, 
			instance.AutoDelete, 
			instance.Exclusive, 
			instance.NoWait, 
			instance.Args,
		)
	}else{
		queue, err = channel.QueueDeclare(
			instance.QueueName, 
			instance.Durable, 
			instance.AutoDelete, 
			instance.Exclusive, 
			instance.NoWait, 
			instance.Args,
		)
	}
	return connection, channel, queue, err
}
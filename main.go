package main

import (
	//"time"

	"github.com/Itsu8/RabbitMQ-communication/errorHandler"
	"github.com/rabbitmq/amqp091-go"
)

func sayHello() string {
	return "Hello, world!"
}

func main() {
	rabbitMQConnection, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	errorHandler.PanicOnError("RabbitMQ server connection error", err)
	defer rabbitMQConnection.Close()

	rabbitMQChannel, err := rabbitMQConnection.Channel()
	errorHandler.PanicOnError("Failed to open a channel", err)
	defer rabbitMQChannel.Close()

	//competing consumers
	err = rabbitMQChannel.Qos(1, 0, false)
	errorHandler.LogOnError("Failed to set QoS", err)

	err = rabbitMQChannel.ExchangeDeclare(
		"testExchange",   
		"fanout", 
		true,     
		false,    
		false,    
		false,    
		nil,      
	)
	errorHandler.PanicOnError("Failed to declare an exchange", err)	
	
	// for{
	// 	time.Sleep(time.Second)
		err = rabbitMQChannel.Publish(
			"testExchange",
			"",
			false,
			false,
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        []byte(sayHello()),
			},
		)
	
		errorHandler.PanicOnError("Failed to send resource", err)
	// }
}

package main

import (
	"fmt"
	"sync"

	"github.com/Itsu8/RabbitMQ-communication/errorHandler"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitMQConnection, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	errorHandler.PanicOnError("RabbitMQ server connection error", err)
	defer rabbitMQConnection.Close()

	rabbitMQChannel, err := rabbitMQConnection.Channel()
	errorHandler.PanicOnError("RabbitMQ channel creation error", err)
	defer rabbitMQChannel.Close()

	rabbitMQueue, err := rabbitMQChannel.QueueDeclare(
		"firstQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	errorHandler.PanicOnError("Failed to create a queue", err)

	err = rabbitMQChannel.QueueBind(
		rabbitMQueue.Name, 
		"",     
		"testExchange", 
		false,
		nil,
	)
	errorHandler.PanicOnError("Failed to bind a queue", err)


	recievedMessages, err := rabbitMQChannel.Consume(
		rabbitMQueue.Name,
		"",
		false, 
		false,
		false,
		false,
		nil,
	)
	errorHandler.PanicOnError("Failed to recieve a message", err)


	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()

	go func(){
		defer wg.Done()
		for msg := range recievedMessages{

			fmt.Println("Recieved messageby reciever №1:", string(msg.Body))
			
			//==========Acknowledging messages by hand==========
			err := msg.Ack(false)
			errorHandler.LogOnError("Failed to acknowledge message",err)

		}
	}()
}

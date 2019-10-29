package main

import(
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	//"math/rand"
	//"time"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf(" %s: %s", msg, err)
	}else{
		log.Fatalf("%s","Connected...")
	}

}
func main() {
	//Dial up the connection...
	conn, err := amqp.Dial(config.AMQPConnectionURL)
	handleError(err,"can not connect")
	defer conn.Close()
	//Create the channel ...
	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")
	defer amqpChannel.Close()
	queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")
	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Could not configure QoS")
	fmt.Printf("%s",queue.Name)
	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not register consumer")
	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)

			addTask := &AddTask{}

			err := json.Unmarshal(d.Body, addTask)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			log.Printf("Result of %d + %d is : %d", addTask.Number1, addTask.Number2, addTask.Number1+addTask.Number2)

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}

		}
	}()

	// Stop for program termination
	<-stopChan

}


package main

import(
	//"fmt"
	"github.com/streadway/amqp"
	"log"
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
	
}
package main

import(
	"fmt"
	"github.com/streadway/amqp"
	"log"
	//"math/rand"
	//"time"
)

func errorHandler(err error, msg string){
	if err != nil {
		log.Fatalf("%s:%s",msg,err)
		fmt.Println("ok now")
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}else{
		log.Fatalf("%s \n","connected")
		fmt.Println("ok now")
	}

}
func main() {
	conn, err := amqp.Dial(config.AMQPConnectionURL)
	//errorHandler(err,"can not connect")
	fmt.Println("testing working")
	fmt.Printf("%s",err)
	errorHandler(err,"can not connect")
	//amqpChannel, err := conn.Channel()
	//errorHandler(err, "Can't create a amqpChannel")
	//queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
	//errorHandler(err, "Could not declare `add` queue")
	//defer amqpChannel.Close()
	defer conn.Close()
}
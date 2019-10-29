package main

import(
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"time"
)

func errorHandler(err error, msg string){
	if err != nil{
		log.Fatalf("%s:%s",msg,err)
	}

}




func main() {
	conn, err := amqp.Dial(shared_conf.Config.AMQPConnectionURL)
}
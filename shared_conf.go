package  main

// Configuration  is used for storing the connection url of rabbitmq 
type Configuration struct {
	AMQPConnectionURL string
}

//AddTask is used for json response here
type AddTask struct {
		Number1 int
		Number2 int
}

var config = Configuration{
	AMQPConnectionURL :"amqp://guest:guest@192.168.56.112:15672/",
}
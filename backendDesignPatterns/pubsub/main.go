package main

// Create two go routines that run infintely , one for publisher
// Exit when the interrupt is received
import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

var connUrl string
var queueName string

func main() {
	connUrl = getAmqURL()
	queueName = "my_q"
	fmt.Println("Url is ", connUrl)
	go publish()
	consume()

}

func failOnError(err error, errmsg string) {
	if err != nil {
		log.Printf("%v \n : %v \n ", errmsg, err.Error())
	}
}

func getAmqURL() string {
	url := os.Getenv("AMQ_URL")
	if url == "" {
		url = "amqp://aldxbzln:gCMk307730CmL_PiUImpFBYwZRvxWVGL@albatross.rmq.cloudamqp.com:5672/aldxbzln"
	}
	return url
}

func publish() {
	conn, err := amqp.Dial(connUrl)
	failOnError(err, "Failed to Connect")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	failOnError(err, "Queue Declare failed")
	for i := 0; i < 10; i++ {
		go func(ch *amqp.Channel) {
			msg := "Hello:" + strconv.Itoa(rand.Int())
			err = ch.Publish("", queueName, false, false, amqp.Publishing{ContentType: "plain/text", Body: []byte(msg)})
			failOnError(err, "Msg Publish Failed")
			fmt.Println("Publish Success: ", msg)
		}(ch)
	}
	time.Sleep(1 * time.Second)
}

func consume() {
	conn, err := amqp.Dial(connUrl)
	failOnError(err, "Failed to Connect")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	failOnError(err, "Queue Declare failed")

	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	failOnError(err, "Failed to consume")

	for msg := range msgs {
		fmt.Printf("Received: %s\n", string(msg.Body))
	}
}

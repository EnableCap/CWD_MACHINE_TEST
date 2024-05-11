package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)


func main() {
	var broker = Broker
	var port = Port
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetDefaultPublishHandler(messagePubHandler)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
  }
	subscribe(client)

	for{
		time.Sleep(time.Second)
	}
	client.Disconnect(250)

}

func subscribe(client mqtt.Client) {
    token := client.Subscribe(Topic, 1, messagePubHandler)
    token.Wait()
    fmt.Printf("Subscribed to topic %s\n", Topic)
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
    fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
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
		opts.OnConnect = connectHandler
		opts.OnConnectionLost = connectLostHandler
		client := mqtt.NewClient(opts)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
	  }
	  	publish(client)
		publishRepeatedly(publish, client)
		client.Disconnect(250)
}


var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
    fmt.Printf("Connect lost: %v", err)
}


func publish(client mqtt.Client) {
	text := fmt.Sprintf("Test Message...")
	token := client.Publish(Topic, 0, false, text)
	token.Wait()
	fmt.Println("Publishing Message")
}

func publishRepeatedly(f func(mqtt.Client), client mqtt.Client){
	for range time.Tick(IntervalInSeconds){
		publish(client)
	}
}

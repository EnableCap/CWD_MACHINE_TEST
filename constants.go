package main
import "time"


const (
	Broker = "localhost"
	Port = 1883
	IntervalInSeconds = 60*time.Second
	Topic = "rooms/room1/test"
)
package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// broker
	broker := "tcp://localhost:1883"
	topic := "test/topic"
	clientID := "do_publisher"

	// MQTT client options
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
	client := mqtt.NewClient(opts)

	// Connect to the broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("接続失敗", token.Error())
		return 
	}
	fmt.Println("接続成功")

	// Publish a message
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Hello, World! %d", i+1)
		token := client.Publish(topic, 0, false, message)
		token.Wait()

		fmt.Println("メッセージ送信: ", message)
		time.Sleep(1 * time.Second)
	}

	// Disconnect from the broker
	client.Disconnect(250)
	fmt.Println("切断")
}
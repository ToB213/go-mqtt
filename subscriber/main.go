package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("受信: トピック [%s] メッセージ: %s\n", msg.Topic(), msg.Payload())
}

func main() {
	// broker
	broker := "tcp://localhost:1883"
	topic := "test/topic"
	clientID := "do_subscriber"

	// MQTT client options
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
	opts.SetDefaultPublishHandler(messagePubHandler)
	client := mqtt.NewClient(opts)

	// Connect to the broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("接続失敗", token.Error())
		return
	}
	fmt.Println("接続成功")

	// Subscribe to the topic
	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println("購読失敗", token.Error())
		return
	}
	fmt.Println("購読成功")

	// Keep the subscriber running
	select {}
}
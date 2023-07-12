package main

import (
	"io"
	"log"
	"os"

	"github.com/rodolfocoding/pubsub/pkg/pubsub"
)

func main() {
	file, err := os.OpenFile("messages.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Configurar o logger para usar uma saída multi-writer que direciona a saída para o arquivo e o terminal
	log.SetOutput(io.MultiWriter(file, os.Stdout))

	broker := pubsub.NewBroker()
	defer broker.Close()

	logger := pubsub.NewMessageLogger("messages.log")

	queue1 := "queue1"
	queue2 := "queue2"

	subscriber1 := broker.Subscribe(queue1)
	subscriber2 := broker.Subscribe(queue2)

	go func() {
		for msg := range subscriber1 {
			log.Printf("Subscriber1 received message: %s\n", msg)
		}
	}()

	go func() {
		for msg := range subscriber2 {
			log.Printf("Subscriber2 received message: %s\n", msg)
		}
	}()

	broker.Publish(pubsub.Message{Queue: queue1, Data: "Message 1 for queue1"})
	broker.Publish(pubsub.Message{Queue: queue2, Data: "Message 1 for queue2"})

	if err := logger.LogMessage(pubsub.Message{Queue: queue1, Data: "Logged message for queue1"}); err != nil {
		log.Fatalf("Failed to log message: %s", err)
	}
}

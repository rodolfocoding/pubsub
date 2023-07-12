package pubsub_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/rodolfocoding/pubsub/pkg/pubsub"
)

func TestPublishSubscribe(t *testing.T) {
	broker := pubsub.NewBroker()
	defer broker.Close()

	queue1 := "queue1"
	queue2 := "queue2"

	subscriber1 := broker.Subscribe(queue1)
	subscriber2 := broker.Subscribe(queue2)

	go func() {
		for msg := range subscriber1 {
			if msg != "Message 1 for queue1" && msg != "Message 2 for queue1" {
				t.Errorf("Unexpected message received by Subscriber1: %s", msg)
			}
		}
	}()

	go func() {
		for msg := range subscriber2 {
			if msg != "Message 1 for queue2" && msg != "Message 2 for queue2" {
				t.Errorf("Unexpected message received by Subscriber2: %s", msg)
			}
		}
	}()

	broker.Publish(pubsub.Message{Queue: queue1, Data: "Message 1 for queue1"})
	broker.Publish(pubsub.Message{Queue: queue2, Data: "Message 1 for queue2"})
	broker.Publish(pubsub.Message{Queue: queue1, Data: "Message 2 for queue1"})
	broker.Publish(pubsub.Message{Queue: queue2, Data: "Message 2 for queue2"})

	time.Sleep(100 * time.Millisecond)
}

func TestMessageLogger(t *testing.T) {
	logFile := "test.log"

	logger := pubsub.NewMessageLogger(logFile)

	queue := "queue1"
	msg := pubsub.Message{Queue: queue, Data: "Logged message"}

	err := logger.LogMessage(msg)
	if err != nil {
		t.Errorf("Failed to log message: %s", err)
	}

	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Errorf("Failed to read log file: %s", err)
	}

	expectedLog := fmt.Sprintf("Queue: %s, Message: %s\n", queue, msg.Data)
	actualLog := string(content)

	if expectedLog != actualLog {
		t.Errorf("Log content does not match. Expected: %q, Actual: %q", expectedLog, actualLog)
	}

	err = os.Remove(logFile)
	if err != nil {
		t.Errorf("Failed to remove log file: %s", err)
	}
}

func TestMain(m *testing.M) {
	result := m.Run()

	os.Exit(result)
}

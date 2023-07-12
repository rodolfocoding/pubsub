package pubsub

import (
	"fmt"
	"os"
	"sync"
)

type MessageLogger struct {
	logFile string
	mutex   sync.Mutex
}

func NewMessageLogger(logFile string) *MessageLogger {
	return &MessageLogger{
		logFile: logFile,
		mutex:   sync.Mutex{},
	}
}

func (l *MessageLogger) LogMessage(msg Message) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	file, err := os.OpenFile(l.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	logLine := fmt.Sprintf("Queue: %s, Message: %s\n", msg.Queue, msg.Data)
	if _, err := file.WriteString(logLine); err != nil {
		return err
	}

	return nil
}

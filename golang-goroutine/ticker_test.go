package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for timeCurrent := range ticker.C {
		fmt.Println(timeCurrent)
	}
}

func TestTic(t *testing.T) {
	channel := time.Tick(1 * time.Second)
	for timeCurrent := range channel {
		fmt.Println(timeCurrent)
	}
}

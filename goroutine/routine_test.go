package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestSelection(t *testing.T) {
	tick := time.Tick(100 * time.Millisecond)
	after := time.After(500 * time.Millisecond)
	for {
		select {
		case a := <-tick:
			fmt.Println("tick.", a)
		case b := <-after:
			fmt.Println("BOOM!", b)
			return
		default:
			fmt.Println("     .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

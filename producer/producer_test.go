package producer

import (
	"strconv"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	for i := 1; ; i++ {
		Send(strconv.Itoa(i))
		time.Sleep(5 * time.Second)
	}

}

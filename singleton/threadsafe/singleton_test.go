package singleton

import (
	"testing"

	"github.com/kaminora/design_pattern/singleton/threadsafe"
)

func TestSingleton(t *testing.T) {
	ch := make(chan interface{})
	go run(ch)
	go run(ch)
	go run(ch)
	<-ch
}

func run(ch chan interface{}) {
	ch <- singleton.GetInstance()
}

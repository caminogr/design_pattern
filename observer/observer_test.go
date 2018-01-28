package observer

import (
	"fmt"
	"testing"

	"github.com/kaminora/design_pattern/observer"
)

func TestObserver(t *testing.T) {
	random := &observer.RandomNumberGenerator{&observer.NumberGenerator{}}

	o1 := &observer.DigitObserver{random}
	o2 := &observer.DigitObserver{random}

	random.AddObserver(o1)
	random.AddObserver(o2)

	result := random.Execute()

	for _, r := range result {
		if len(result) != 2 && r >= 50 {
			t.Errorf("Expect result to equal random int array")
		}
	}
}

package singleton

import (
	"testing"

	"github.com/kaminora/design_pattern/singleton"
)

func TestSingleton(t *testing.T) {

	t.Log("Start")
	singleton1 := singleton.GetInstance()
	singleton2 := singleton.GetInstance()

	if singleton1 == singleton2 {
		t.Log("2つは同じインスタンスです")
	} else {
		t.Log("2つは同じインスタンスではありません")
	}
}

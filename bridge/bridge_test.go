package bridge

import (
	"testing"

	"github.com/kaminora/design_pattern/bridge"
)

func TestBridge(t *testing.T) {
	d1 := bridge.DefaultDisplay{&bridge.StringDisplayImpl{"AAA"}}
	expect := "+---+\n|AAA|\n+---+\n"
	if result := d1.Display(); result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}

	d2 := bridge.CountDisplay{&bridge.DefaultDisplay{&bridge.StringDisplayImpl{"BBB"}}}
	expect = "+---+\n|BBB|\n+---+\n"
	if result := d2.DefaultDisplay.Display(); result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}

	expect = "+---+\n|BBB|\n|BBB|\n+---+\n"
	if result := d2.MultiDisplay(2); result != expect {
		t.Errorf("Expect result to equal \n%s\n, but \n%s.\n", expect, result)
	}
}

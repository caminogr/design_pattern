package facade

import (
	"testing"

	"github.com/kaminora/design_pattern/facade"
)

func TestFacade(t *testing.T) {

	maker := facade.PageMaker{}
	result := maker.MakeWelcomePage("hoge@example.com")
	expect := "# Welcome to hoge`s page!!"

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}

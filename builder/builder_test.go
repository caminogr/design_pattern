package builder

import (
	"testing"

	"github.com/kaminora/design_pattern/builder"
)

func TestBuilder(t *testing.T) {
	expect := "<header>Header</header>\n" +
		"<article>Content</article>\n" +
		"<footer>Footer</footer>\n"

	product := &builder.Product{}

	director := builder.Director{&builder.HTMLBuilder{product}}
	director.Construct()

	result := product.Show()

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s", expect, result)
	}
}

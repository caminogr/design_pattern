package builder

// Builder
type Builder interface {
	makeHeader(str string) string
	makeContent(str string) string
	makeFooter(str string) string
}

// Director
type Director struct {
	Builder
}

func (self *Director) Construct() string {
	return self.Builder.makeHeader("Header") +
		self.Builder.makeContent("Content") +
		self.Builder.makeFooter("Footer")
}

// ConcreteBuiler
type HTMLBuilder struct {
	*Product
}

func (self *HTMLBuilder) makeHeader(str string) string {
	self.Product.Header = "<header>" + str + "</header>\n"
	return self.Product.Header
}

func (self *HTMLBuilder) makeContent(str string) string {
	self.Product.Content = "<article>" + str + "</article>\n"
	return self.Product.Content
}

func (self *HTMLBuilder) makeFooter(str string) string {
	self.Product.Footer = "<footer>" + str + "</footer>\n"
	return self.Product.Footer
}

type Product struct {
	Header  string
	Content string
	Footer  string
}

func (self *Product) Show() string {
	var result string
	result = self.Header
	result += self.Content
	result += self.Footer
	return result
}

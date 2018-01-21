package bridge

type DisplayImpl interface {
	rawOpen() string
	rawPrint() string
	rawClose() string
}

type StringDisplayImpl struct {
	Str string
}

func (self *StringDisplayImpl) rawOpen() string {
	return self.printLine()
}

func (self *StringDisplayImpl) rawPrint() string {
	return "|" + self.Str + "|\n"
}

func (self *StringDisplayImpl) rawClose() string {
	return self.printLine()
}

func (self *StringDisplayImpl) printLine() string {
	Str := "+"
	for _, _ = range self.Str {
		Str += string("-")
	}
	Str += "+\n"
	return Str
}

type DefaultDisplay struct {
	Impl DisplayImpl
}

func (self *DefaultDisplay) Open() string {
	return self.Impl.rawOpen()
}

func (self *DefaultDisplay) Print() string {
	return self.Impl.rawPrint()
}

func (self *DefaultDisplay) Close() string {
	return self.Impl.rawClose()
}

func (self *DefaultDisplay) Display() string {
	return self.Open() +
		self.Print() +
		self.Close()
}

type CountDisplay struct {
	*DefaultDisplay
}

func (self *CountDisplay) MultiDisplay(times int) string {
	Str := self.Open()
	for i := 0; i < times; i++ {
		Str += self.Print()
	}
	Str += self.Close()
	return Str
}

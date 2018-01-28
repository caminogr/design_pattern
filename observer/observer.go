package observer

import (
	"math/rand"
)

// Subject
type NumberGenerator struct {
	observers []Observer
}

func (self *NumberGenerator) AddObserver(observer Observer) {
	self.observers = append(self.observers, observer)
}

func (self *NumberGenerator) NotifyObservers() []int {
	var result []int
	for _, observer := range self.observers {
		result = append(result, observer.Update())
	}
	return result
}

// ConcreteSubject
type RandomNumberGenerator struct {
	*NumberGenerator
}

func (self *RandomNumberGenerator) GetNumber() int {
	return rand.Intn(50)
}

func (self *RandomNumberGenerator) Execute() []int {
	return self.NotifyObservers()
}

// Observer
type Observer interface {
	Update() int
}

type number interface {
	GetNumber() int
}

// ConcreteObserver
type DigitObserver struct {
	Generator number
}

func (self *DigitObserver) Update() int {
	return self.Generator.GetNumber()
}

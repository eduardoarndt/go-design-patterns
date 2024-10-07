package main

import (
	"fmt"
	"sync"
)

type Singleton interface {
	DoSomething() string
}

type singleton struct{}

var lock = &sync.Mutex{}

var instance *singleton

func NewSingletonInstance() *singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating single instance now.")
			instance = &singleton{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return instance
}

func (s *singleton) DoSomething() string {
	return "Doing something."
}

func main() {
	instance1 := NewSingletonInstance()

	instance2 := NewSingletonInstance()

	fmt.Printf("%p\n", instance1)

	fmt.Printf("%p\n", instance2)
}

# Singleton

The Singleton pattern ensures that a class has only one instance and provides a global point of access to that instance.

This is useful in cases where you need to manage a shared resource, such as database connections or configuration settings.

### Problem Statement

Often there is the need to ensure that only one instance of a class exists, such as when managing configurations or interacting with hardware resources. Without Singleton, creating multiple instances can lead to issues like inconsistent data or resource locks.

This is very common to happen when working with asynchronous code, where multiple goroutines can create new instances of a class or access shared resources.

### Real-World Example

Consider a database connection pool: If multiple parts of your application create new connections at the same time, you may end up with redundant or conflicting database interactions. A Singleton ensures only one connection is created and used across the application.

### Implementation

- Run the example by running `go run singleton/main.go`.

```go
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
```

#### Explanation

The function `NewSingletonInstance` ensures that only one instance of singleton is created, even when called multiple times.

- First, it checks if instance is nil (i.e., no instance has been created yet).
- If instance is nil, it locks the section of code using lock.Lock() to prevent multiple Goroutines from entering this section simultaneously.
- After locking, a second check (if instance == nil) is performed to ensure that no other Goroutine created the instance between the first check and the time the lock was acquired.
- If the instance is still nil, a new singleton instance is created and assigned to the global instance variable.
- The sync.Mutex and double-checked locking ensure that the creation of the singleton instance is thread-safe, preventing multiple Goroutines from creating separate instances.

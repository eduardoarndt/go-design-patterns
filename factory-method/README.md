## Factory Method

The Factory Method is a creational design pattern that provides an interface for creating objects but allows subclasses to alter the type of objects that will be created. Instead of directly instantiating objects using new, the Factory Method delegates the responsibility of object creation to subclasses or methods, promoting flexibility and scalability.

### Problem Statement

When there's the need to create objects, but you don’t know in advance the exact class of the object that will be required. The application needs to handle different types of objects that share common behavior but differ in their implementation. Additionally, you want to avoid modifying your code each time a new type of object is added or when the way objects are created changes.

### Specific Problems It Solves

- Complex object creation: Directly instantiating classes can tie your code to specific implementations, making it rigid and difficult to maintain.

- Dynamic object creation: In many cases, the type of object required may not be known until runtime, and this decision should be flexible.

- Encapsulation of instantiation: Object creation logic should be separated from the client code, so changes in how objects are created don't affect the rest of the system.

- Scalability: When you need to add new types of products (objects), you want a scalable solution that avoids modifying existing code.

### Real-World Example

Imagine a car factory where you place an order for a car, but depending on your preferences (e.g., electric or gas), the factory assembles and provides you with the appropriate car type. The client doesn't need to know the specific details of car assembly—they just receive the product.

### Implementation

```go
package main

import (
	"fmt"
	"os"
)

// Step 1: Define the Product Interface
type Car interface {
	Drive() string
	FuelType() string
}

// Step 2: Concrete Products (Electric Car and Gas Car)
type ElectricCar struct{}

func (e *ElectricCar) Drive() string {
	return "Driving an electric car"
}
func (e *ElectricCar) FuelType() string {
	return "Powered by electricity"
}

type GasCar struct{}

func (g *GasCar) Drive() string {
	return "Driving a gas-powered car"
}

func (g *GasCar) FuelType() string {
	return "Powered by gasoline"
}

// Step 3: Define the Factory Interface
type CarFactory interface {
	CreateCar(brand string) Car
}

type carFactory struct{}

func (carFactory *carFactory) CreateCar() Car {
	carPreference := os.Getenv("CAR_PREFERENCE")

	if carPreference == "electric" {
		return &ElectricCar{}
	}

	return &GasCar{}
}

// Step 4: Client Code
func main() {
	// Client uses the factory to create objects
	carFactory := carFactory{}

	// Creating a Gas Car
	gasCar := carFactory.CreateCar()
	fmt.Println(gasCar.Drive())    // Output: Driving a gas-powered car
	fmt.Println(gasCar.FuelType()) // Output: Powered by gasoline

	// Creating an Electric Car
	os.Setenv("CAR_PREFERENCE", "electric")
	electricCar := carFactory.CreateCar()
	fmt.Println(electricCar.Drive())    // Output: Driving an electric car
	fmt.Println(electricCar.FuelType()) // Output: Powered by electricity
}
```

Product Interface: Defines a common interface (Car) with methods Drive() and FuelType().

Concrete Products: Implements the interface with specific classes (ElectricCar, GasCar) that define their behavior.

Factory Interface: Specifies a method (CreateCar()) for creating Car objects.

Factory Logic: The factory decides which type of car to create based on the client's preferences. This is exampled here by using environment variables, but it could be based on any logic.

Client Code: Requests cars from the factory without knowing the creation details, using the returned objects through the common interface.

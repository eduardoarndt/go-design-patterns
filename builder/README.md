# Builder

The Builder Pattern is a creational design pattern that provides a flexible solution for constructing complex objects step by step. It allows you to create different representations of an object using the same construction process.

### Problem Statement

- When creating an object involves multiple steps or parameters, managing the construction process can become cumbersome.
- You want to avoid a "telescoping constructor" problem (where constructors take many parameters) and ensure that your object is immutable after creation.

### Benefits of the Builder Pattern

- Flexibility: Easily construct different representations of an object.
- Readability: Method chaining makes the construction process clear and concise.
- Encapsulation: The construction logic is encapsulated within the builder, separating it from the client code.

### Real-World Example

Consider building a computer. A computer can have various components like CPU, RAM, storage, etc.

Instead of requiring all these parameters in a constructor, the builder pattern allows you to set each component step by step.

### Implementation

```go
package main

import "fmt"

// Product
type Computer struct {
	CPU     string
	RAM     string
	Storage string
	OS      string
}

// Builder Interface
type ComputerBuilder interface {
	SetCPU(cpu string) ComputerBuilder
	SetRAM(ram string) ComputerBuilder
	SetStorage(storage string) ComputerBuilder
	SetOS(os string) ComputerBuilder
	Build() *Computer
}

// Concrete Builder
type PCBuilder struct {
	computer *Computer
}

func NewPCBuilder() *PCBuilder {
	return &PCBuilder{&Computer{}}
}

func (b *PCBuilder) SetCPU(cpu string) ComputerBuilder {
	b.computer.CPU = cpu
	return b
}

func (b *PCBuilder) SetRAM(ram string) ComputerBuilder {
	b.computer.RAM = ram
	return b
}

func (b *PCBuilder) SetStorage(storage string) ComputerBuilder {
	b.computer.Storage = storage
	return b
}

func (b *PCBuilder) SetOS(os string) ComputerBuilder {
	b.computer.OS = os
	return b
}

func (b *PCBuilder) Build() *Computer {
	return b.computer
}

// Client Code
func main() {
	builder := NewPCBuilder()

	computer := builder.
		SetCPU("Intel i7").
		SetRAM("16GB").
		SetStorage("1TB SSD").
		SetOS("Windows 7").
		Build()

	fmt.Printf("Computer built: %+v\n", computer)
}
```

Product (Computer): Represents the complex object that will be built with various attributes.

Builder Interface (ComputerBuilder): Defines methods for setting each component of the product.

Concrete Builder (PCBuilder): Implements the builder interface.

- Provides methods to set CPU, RAM, Storage, and OS, returning the builder itself for method chaining.
- The Build() method returns the final Computer object.

Client Code: Creates a new builder instance.

- Configures the computer step by step using method chaining.
- Calls Build() to retrieve the final product.

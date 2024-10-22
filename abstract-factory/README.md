## Abstract Factory

Abstract Factory is a creational design pattern that lets you create related objects without specifying their concrete classes.

### Problem Statement

Imagine you are developing a GUI toolkit that should support multiple look-and-feel standards (e.g., Windows, macOS, Linux). Each look-and-feel standard requires a different set of UI components such as buttons, checkboxes, and scrollbars.

Without the Abstract Factory pattern, you would have to write a lot of conditional code to handle the creation of these components based on the current look-and-feel standard. This approach is not scalable and makes the codebase difficult to maintain and extend.

By using the Abstract Factory pattern, you can create an interface for creating families of related objects (e.g., buttons, checkboxes, scrollbars) and implement this interface for each look-and-feel standard. This way, the client code can work with any look-and-feel standard without knowing the specifics of how the components are created.

### Real-World Example

Consider a furniture factory that produces chairs and tables in different styles such as classic and modern.

Each style requires a different set of components (e.g., legs, backrest, armrests) and materials (e.g., wood, metal, plastic).

### Implementation

- Run the example by running `go run abstract-factory/main.go`.

```go
package main

import "fmt"

// Abstract Factory
type FurnitureFactory interface {
	MakeChair() Chair
	MakeTable() Table
}

func NewFurnitureFactory(style string) (FurnitureFactory, error) {
	if style == "classic" {
		return &Classic{}, nil
	}

	if style == "modern" {
		return &Modern{}, nil
	}

	return nil, fmt.Errorf("Style %s is not supported", style)
}

type Chair struct {
	Name     string
	Material string
}

type Table struct {
	Name     string
	Material string
}

// Concrete Classic Factory
type Classic struct{}

func (a *Classic) MakeChair() Chair {
	return Chair{
		Name:     "Classic Chair",
		Material: "Wood",
	}
}

func (a *Classic) MakeTable() Table {
	return Table{
		Name:     "Classic Table",
		Material: "Wood",
	}
}

// Concrete Modern Factory
type Modern struct{}

func (n *Modern) MakeChair() Chair {
	return Chair{
		Name:     "Modern Chair",
		Material: "Plastic",
	}
}

func (n *Modern) MakeTable() Table {
	return Table{
		Name:     "Modern Table",
		Material: "Plastic",
	}
}

func main() {
	factory, _ := NewFurnitureFactory("classic")

	chair := factory.MakeChair()
	table := factory.MakeTable()

	fmt.Printf("Chair: %s\n", chair.Name)
	fmt.Printf("Table: %s\n", table.Name)

	factory, _ = NewFurnitureFactory("modern")

	chair = factory.MakeChair()
	table = factory.MakeTable()

	fmt.Printf("Chair: %s\n", chair.Name)
	fmt.Printf("Table: %s\n", table.Name)
}

```

#### Explanation

This code demonstrates the Abstract Factory pattern by creating furniture (chairs and tables) in different styles (classic and modern). It defines an interface for creating families of related objects and concrete implementations for each style, allowing easy scalability and maintenance.

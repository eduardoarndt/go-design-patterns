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

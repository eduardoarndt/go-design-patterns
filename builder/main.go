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

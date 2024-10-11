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

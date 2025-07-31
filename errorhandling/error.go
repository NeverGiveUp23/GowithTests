package main

import (
	"errors"
	"fmt"
	"log"
)

// for more design you can make err variable to be flexible.
var (
	ErrNotImplemented = errors.New("truck not processed")
	ErrTruckNotFound  = errors.New("truck not found")
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type GasTruck struct {
	id    string
	cargo int
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (t *GasTruck) LoadCargo() error {
	return ErrTruckNotFound
}

func (e *ElectricTruck) LoadCargo() error {
	return ErrTruckNotFound
}

func (e *ElectricTruck) UnloadCargo() error {
	return ErrTruckNotFound
}

func (t *GasTruck) UnloadCargo() error {
	return ErrTruckNotFound
}

func processTrucks(truck Truck) error {
	fmt.Printf("processing truck: %s\n", truck)
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}
	return ErrNotImplemented
}

func main() {

	trucks := []GasTruck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	//eTrucks := []ElectricTruck{
	//	{id: "eTruck-1"},
	//	{id: "eTruck-2"},
	//	{id: "eTruck-3"},
	//}

	err := processTrucks(&GasTruck{id: "1"})
	if err != nil {
		log.Fatalf("Error processing truck: %s", err)
	}
	err = processTrucks(&ElectricTruck{id: "2"})
	if err != nil {
		log.Fatalf("Error processing truck: %s", err)
	}

}

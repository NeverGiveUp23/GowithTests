package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

// for more design you can make err variable to be flexible.
var (
	ErrNotImplemented = errors.New("truck not processed")
	ErrTruckNotFound  = errors.New("truck not found")
)

var (
	_, SuccessToLoadCargo = fmt.Println("Truck is ready to load")
)

type contextKey string

var userIDKey contextKey = "userID"

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

// If you want to make a generic map you can create a function, this is option 1
func makeMap[K comparable, V any]() map[K]V {
	return make(map[K]V)
}

// option 2 type alias from Go 1.18+
type GenericMap[K comparable, V any] map[K]V

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
	t.cargo += 1
	return nil
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery -= 1
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo -= 1
	e.battery -= 1
	return ErrTruckNotFound
}

func (t *GasTruck) UnloadCargo() error {
	t.cargo = 0
	return ErrTruckNotFound
}

func processTrucks(ctx context.Context, truck Truck) error {
	fmt.Printf("started processing truck: %+v \n", truck)

	//userID := ctx.Value(userIDKey)
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	delay := time.Second * 3
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		break
	}

	err := truck.LoadCargo()
	if err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	err = truck.UnloadCargo()
	if err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	fmt.Printf("finished processing truck: %+v \n", truck)
	return nil
}

func processFleet(ctx context.Context, trucks []Truck) error {
	var wg sync.WaitGroup
	errorsChan := make(chan error)

	defer close(errorsChan)

	for _, t := range trucks {
		wg.Add(1)

		go func(t Truck) {
			if err := processTrucks(ctx, t); err != nil {
				errorsChan <- err // send error to the channel
			}

			wg.Done()
		}(t)
	}
	wg.Wait()

	select {
	case err := <-errorsChan:
		return err
	default:
		return nil
	}
}

func main() {

	ctx := context.Background() // carry info and cancel to control the flow
	ctx = context.WithValue(ctx, userIDKey, 42)

	// Option 1: Function
	person := makeMap[string, interface{}]()
	person["name"] = "Felix"
	person["age"] = 32

	age, exist := person["age"].(int)
	if !exist {
		log.Fatal("age does not exist")
		return
	}
	log.Println(age)

	// Option 2: Type Alias
	person2 := GenericMap[string, interface{}]{}
	person2["name"] = "Generic"

	fleet := []Truck{
		&GasTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&GasTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	if err := processFleet(ctx, fleet); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
		return
	}

	fmt.Println("All trucks processed")

}

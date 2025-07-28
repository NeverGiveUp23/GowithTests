package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	// orderChan := make(chan *Order) // unbuffered cannels
	buffOrderedChan := make(chan *Order, 20) // buffered channel -> 20 values inside

	go func() {
		defer wg.Done()
		for _, order := range generateOrders(20) {
			buffOrderedChan <- order
		}

		close(buffOrderedChan)

		fmt.Println("Done with generating orders")
	}()

	go processOrders(buffOrderedChan, &wg)

	// workers
	//for i := 0; i < 3; i++ {
	//	go func() {
	//		defer wg.Done()
	//		for _, order := range orders {
	//			updateOrderStatus(order)
	//		}
	//	}()
	//}

	wg.Wait()

	fmt.Println("All operation completed")
	// fmt.Println(totalUpdates)
}

func updateOrderStatus(order *Order) {
	order.mu.Lock()
	time.Sleep(
		time.Duration(rand.Intn(300)) * time.Millisecond,
	)
	status := []string{
		"Processing", "Shipped", "Delivered",
	}[rand.Intn(3)]
	order.Status = status
	fmt.Printf("Updated order %d status: %s\n", order.ID, status)
	order.mu.Unlock()

	updateMutex.Lock()
	defer updateMutex.Unlock() // unlock mutex
	currentUpdates := totalUpdates
	time.Sleep(5 * time.Millisecond)
	totalUpdates = currentUpdates + 1
}

func processOrders(orderChan <-chan *Order, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orderChan {
		time.Sleep(time.Duration(
			rand.Intn(500)) *
			time.Millisecond,
		)
		fmt.Printf("Processing order %d\n", order.ID)
	}
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{ID: i + 1, Status: "Pending"}
	}

	return orders
}

func reportOrderStatus(orders []*Order) {
	fmt.Println("\n--- Order Status Report ---")
	for _, order := range orders {
		fmt.Printf(
			"Order %d: %s\n", order.ID, order.Status,
		)
	}
	fmt.Println("--------------------------\n")
}

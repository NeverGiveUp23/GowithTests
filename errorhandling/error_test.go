package main

import (
	"testing"
)

func TestLoad(t *testing.T) {
	t.Run("process truck", func(t *testing.T) {
		t.Run("should load and unload cargo", func(t *testing.T) {
			nt := &GasTruck{id: "1"}
			et := &ElectricTruck{id: "2"}

			err := processTrucks(nt)
			if err != nil {
				t.Fatalf("Error processing truck: %s", err)
			}
			err = processTrucks(et)
			if err != nil {
				t.Fatalf("Error processing truck: %s", err)
			}

			// asserting
			if nt.cargo != 0 {
				t.Fatal("normal cargo should be zero")
			}
			if et.battery != -1 {
				t.Fatal("Battery should be minus 1")
			}
		})
	})
}

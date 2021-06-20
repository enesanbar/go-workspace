package factory

import (
	"testing"

	"github.com/enesanbar/workspace/golang/design-patterns/creational/04-abstract-factory/car"
)

func TestCarFactory_GetVehicle(t *testing.T) {
	carF, err := GetVehicleFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carF.GetVehicle(car.LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Car vehicle has %d seats\n", carVehicle.GetWheels())

	luxuryCar, ok := carVehicle.(car.Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Luxury car has %d doors.\n", luxuryCar.GetDoors())
}

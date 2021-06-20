package factory

import (
	"testing"

	"github.com/enesanbar/workspace/golang/design-patterns/creational/04-abstract-factory/motorbike"
)

func TestMotorbikeFactory_GetVehicle(t *testing.T) {
	motorbikeF, err := GetVehicleFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeF.GetVehicle(motorbike.SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Motorbike vehicle has %d wheels\n", motorbikeVehicle.GetWheels())

	sportBike, ok := motorbikeVehicle.(motorbike.Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Sport motorbike has type %d\n", sportBike.GetType())
}

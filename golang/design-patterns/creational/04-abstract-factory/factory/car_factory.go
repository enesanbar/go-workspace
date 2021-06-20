package factory

import (
	"fmt"

	abstractFactory "github.com/enesanbar/workspace/golang/design-patterns/creational/04-abstract-factory"
	"github.com/enesanbar/workspace/golang/design-patterns/creational/04-abstract-factory/car"
)

type CarFactory struct{}

func (c *CarFactory) GetVehicle(v int) (abstractFactory.Vehicle, error) {
	switch v {
	case car.LuxuryCarType:
		return new(car.LuxuryCar), nil
	case car.FamilyCarType:
		return new(car.FamilyCar), nil
	default:
		return nil, fmt.Errorf("Vehicle of type %d not recognized\n", v)
	}
}

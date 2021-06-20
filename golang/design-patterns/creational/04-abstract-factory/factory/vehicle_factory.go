package factory

import (
	"fmt"

	abstractFactory "github.com/enesanbar/workspace/golang/design-patterns/creational/04-abstract-factory"
)

type VehicleFactory interface {
	GetVehicle(v int) (abstractFactory.Vehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, fmt.Errorf("Factory with id %d not recognized\n", f)
	}
}

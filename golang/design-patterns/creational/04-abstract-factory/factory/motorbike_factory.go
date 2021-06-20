package factory

import (
	"fmt"

	abstractFactory "github.com/enesanbar/workspace/golang/design-patterns/creational/04-abstract-factory"
	"github.com/enesanbar/workspace/golang/design-patterns/creational/04-abstract-factory/motorbike"
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) GetVehicle(v int) (abstractFactory.Vehicle, error) {
	switch v {
	case motorbike.SportMotorbikeType:
		return new(motorbike.SportMotorbike), nil
	case motorbike.CruiseMotorbikeType:
		return new(motorbike.CruiseMotorbike), nil
	default:
		return nil, fmt.Errorf("Vehicle of type %d not recognized\n", v)
	}
}

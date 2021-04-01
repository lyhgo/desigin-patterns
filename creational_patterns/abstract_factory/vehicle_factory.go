package abstract_factory

import (
	"fmt"
)
// When to use
// 1. Provide a new layer of encapsulation for Factory methods that return a common
// interface for all factories.
// 2. Group common factories into a super Factory (also called a factory of factories).
type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory),nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory),nil
	default:
		return nil, fmt.Errorf("factory with id %d notrecognized", f)
	}
}

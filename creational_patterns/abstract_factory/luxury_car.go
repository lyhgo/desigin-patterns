package abstract_factory

type LuxuryCar struct{}

func (f *LuxuryCar) NumDoors() int {
	return 4
}

func (f *LuxuryCar) NumWheels() int {
	return 4
}

func (f *LuxuryCar) NumSeats() int {
	return 5
}

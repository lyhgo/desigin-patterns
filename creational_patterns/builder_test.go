package creational_patterns

import "testing"

func TestBuilder(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("excepted wheel of car is 4, but get %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("excepted get 'Car' structrue with car builder, but get %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("excepted seats of car is 5, but get %d\n", car.Seats)
	}

	motorbikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(motorbikeBuilder)
	manufacturingComplex.Construct()

	motorbike := motorbikeBuilder.GetVehicle()

	if motorbike.Wheels != 2 {
		t.Errorf("excepted wheel of motorbike is 4, but get %d\n", motorbike.Wheels)
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("excepted get 'Motorbike' structrue with motorbike builder, but get %s\n", motorbike.Structure)
	}

	if motorbike.Seats != 2 {
		t.Errorf("excepted seats of motorbike is 5, but get %d\n", motorbike.Seats)
	}
}

package abstract_factory

import "testing"

func TestCarFactory(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatalf("Failed to get car factory, error message:\n%s", err.Error())
	}
	familyCarVehicle, err := carF.Build(FamilyCarType)
	if err != nil {
		t.Fatalf("Failed to build family car, error message:\n%s", err.Error())
	}

	seats := familyCarVehicle.NumSeats()
	wheels := familyCarVehicle.NumWheels()
	if seats != 5 || wheels != 4 {
		t.Fatalf("The family car should have 5 seats and 4 wheels, but get %d seats,%d wheels\n", seats, wheels)
	}
	familyCar := familyCarVehicle.(Car)
	if doors := familyCar.NumDoors(); doors != 5 {
		t.Fatalf("The family car should have 5 doors, but get %d\n", doors)
	}

	luxuryCarVehicle, err := carF.Build(LuxuryCarType)

	seats = luxuryCarVehicle.NumSeats()
	wheels = luxuryCarVehicle.NumWheels()
	if seats != 5 || wheels != 4 {
		t.Fatalf("The luxury car should have 5 seats and 4 wheels, but get %d seats,%d wheels\n", seats, wheels)
	}

	luxuryCar := luxuryCarVehicle.(Car)
	if err != nil {
		t.Fatalf("Failed to build luxury car, error message:\n%s", err.Error())
	}
	if doors := luxuryCar.NumDoors(); doors != 4 {
		t.Fatalf("The luxury car excepted get 4 doors, but get %d\n", doors)
	}

}

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatalf("Failed to get motorbike factory, error message:\n%s", err.Error())
	}
	cruiseMotorbikeVehicle, err := motorbikeF.Build(CruiseMotorbikeType)
	if err != nil {
		t.Fatalf("Failed to build cruise motorbike, error message:\n%s", err.Error())
	}

	seats := cruiseMotorbikeVehicle.NumSeats()
	wheels := cruiseMotorbikeVehicle.NumWheels()
	if seats != 1 || wheels != 2 {
		t.Fatalf("The cruise motorbike should have 1 seats and 2 wheels, but get %d seats,%d wheels\n", seats, wheels)
	}
	cruiseMotorbike := cruiseMotorbikeVehicle.(Motorbike)
	if code := cruiseMotorbike.GetMotorbikeType(); code != 2 {
		t.Fatalf("The code of cruise motorbike should be 2, but get %d\n", code)
	}

	sportMotorbikeVehicle, err := motorbikeF.Build(SportMotorbikeType)

	seats = sportMotorbikeVehicle.NumSeats()
	wheels = sportMotorbikeVehicle.NumWheels()
	if seats != 1 || wheels != 2 {
		t.Fatalf("The sport motorbike should have 1 seats and 2 wheels, but get %d seats,%d wheels\n", seats, wheels)
	}

	sportMotorbike := sportMotorbikeVehicle.(Motorbike)
	if err != nil {
		t.Fatalf("Failed to build luxury car, error message:\n%s", err.Error())
	}
	if code := sportMotorbike.GetMotorbikeType(); code != 1 {
		t.Fatalf("The code of sport motorbike should be 1, but get %d\n", code)
	}
}

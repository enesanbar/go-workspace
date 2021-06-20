package car

type FamilyCar struct{}

func (l *FamilyCar) GetDoors() int {
	return 5
}

func (l *FamilyCar) GetWheels() int {
	return 4
}

func (l *FamilyCar) GetSeats() int {
	return 5
}

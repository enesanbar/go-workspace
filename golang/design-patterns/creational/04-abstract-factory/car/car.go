package car

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type Car interface {
	GetDoors() int
}

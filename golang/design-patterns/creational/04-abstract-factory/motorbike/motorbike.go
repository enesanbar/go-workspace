package motorbike

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type Motorbike interface {
	GetType() int
}

package helpers

type Runnable interface {
	Start() error
	Stop() error
}

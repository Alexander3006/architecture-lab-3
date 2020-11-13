package interfaces

type HttpPortNumber int

type IServer interface {
	Start(HttpPortNumber) error
}

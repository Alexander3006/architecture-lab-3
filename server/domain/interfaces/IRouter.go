package interfaces

type IRouter interface {
	RegisterRoute(string, string, IController) error
	IndicateRoute(IConnection)
	CreateController(func(IConnection)) IController
}
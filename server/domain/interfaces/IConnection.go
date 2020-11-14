package interfaces

import "net/http"

type IConnection interface {
	GetRequest() *http.Request
	SendError(int, string)
	Ok(interface{})
	NotFound(string)
	GetBody(interface{}) error
}




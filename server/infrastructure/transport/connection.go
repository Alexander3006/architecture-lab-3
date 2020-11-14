package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Connection struct{
	request *http.Request
	response http.ResponseWriter
}

func (c Connection) writeJson(status int, res interface{}) {
	rw := c.response
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(status)
	err := json.NewEncoder(rw).Encode(res)
	if err != nil {
		fmt.Printf("Error writing response: %s", err)
	}
}

func (c Connection) GetRequest() *http.Request {
	return c.request
}

func (c Connection) SendError(code int, message string) {
	c.writeJson(code, message)
}

func (c Connection) NotFound(message string){
	c.writeJson(http.StatusNotFound, message)
}

func (c Connection) Ok(res interface{}) {
	c.writeJson(http.StatusOK, res)
}

func (c Connection) GetBody(buffer interface{}) error {
	err := json.NewDecoder(c.request.Body).Decode(buffer)
	if err != nil {
		c.SendError(400, "bad Json payload")
	}
	return err
}
package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

func main() {
	err := server.StartServer(New, Version, Priority)
	if err != nil {
		return
	}
}

var Version = "1.0"
var Priority = 1

type Config struct {
	RequestHeader  string `json:"request_header"`
	ResponseHeader string `json:"response_header"`
}

func New() interface{} { return &Config{} }

// Executa antes de realizar a request

func (conf Config) Access(kong *pdk.PDK) {
	err := kong.ServiceRequest.SetHeader("X-TigreDoMexico-Header", conf.RequestHeader)
	if err != nil {
		return
	}
}

// Executa antes de retornar o response

func (conf Config) Response(kong *pdk.PDK) {
	err := kong.Response.SetHeader("X-TigreDoMexico-Header", conf.ResponseHeader)
	if err != nil {
		return
	}
}

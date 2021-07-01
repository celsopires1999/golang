package main

import (
	"github.com/celsopires1999/golang/data"
	"github.com/celsopires1999/golang/webserver"
)

func main ()  {
	
	data.LoadData()
	webserver := webserver.NewWebServer()
	webserver.Serve()


}

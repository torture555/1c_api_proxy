package main

import (
	"1c_api_proxy/internal/app"
	"flag"
)

func main() {

	var port string

	flag.StringVar(&port, "port", "", "Enter port service")
	flag.Parse()

	app.StartConnectionSQL()

	app.StartLogger()

	app.StartServices1CAPI()

	app.StartRouteProxy(port)

}

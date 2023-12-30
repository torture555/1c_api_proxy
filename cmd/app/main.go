package main

import (
	"1c_api_proxy/internal/app"
)

func main() {

	app.StartConnectionSQL()

	app.StartLogger()

	app.StartServices1CAPI()

	app.StartRouteProxy()

}

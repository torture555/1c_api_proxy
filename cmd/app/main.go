package main

import (
	"1c_api_proxy/internal/app"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	app.StartConnectionSQL()

	app.StartFileLogger()

	app.StartServices1CAPI()

	wg.Add(1)
	go app.StartRouteProxy(&wg)

	wg.Add(1)
	go app.StartBackend(&wg)

	wg.Add(1)
	go app.StartFrontEnd(&wg)

	wg.Wait()

}

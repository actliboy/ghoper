package main

import (
	"github.com/kataras/iris"
	"log"
	"net/http"
	"service/controller/crawler"
	"service/controller/cron"
	"service/controller/hwebsocket"
	"service/initialize"
	"service/router"
)

func main() {

	crawler.M131()
	//go crawler.MM131()

	cron.New().Start()
	defer cron.New().Stop()

	go hwebsocket.Start()

	//go hcache.Start()

	irisRouter := router.IrisRouter()

	// listen and serve on http://0.0.0.0:8080.
	if err := irisRouter.Run(iris.Addr(initialize.Config.Server.HttpPort),
		iris.WithConfiguration(iris.YAML("../config/iris.yml"))); err != nil && err != http.ErrServerClosed {
		log.Printf("Listen: %s\n", err)
	}

	/*	opts := groupcache.HTTPPoolOptions{BasePath: hcache.BasePath}
		peers := groupcache.NewHTTPPoolOpts("", &opts)
		peers.Set("http://localhost:8333", "http://localhost:8222")

		val, err := hcache.GetFromPeer("helloworld", "wjs1", peers)*/

}

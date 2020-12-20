package main

import (
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/markbates/pkger"
	"github.com/phayes/freeport"
	"go-universal-interface/server"
	"log"
	"net/http"
	"strconv"
)

func main() {
	devMode := flag.Bool("devMode", false, "Use Live Relod Dev Server")
	port := flag.Int("port", 8000, "Port to serve app on")
	host := flag.String("host", "localhost", "Hostname to serve on")
	flag.Parse()

	router := httprouter.New()
	counter := server.Counter{
		Value: 0,
	}

	server.AddCounterApiRoutes(router, "/api", &counter)
	pkger.Include("/ui/build")

	if *devMode {
		devServerPort, err := freeport.GetFreePort()
		if err != nil {
			log.Fatal(err)
		}
		server.StartUIDevServer("/ui", devServerPort)
		server.AddUIProxyRoutes(router, "/ui", "localhost", devServerPort)
	} else {
		server.AddUiStaticServer(router, "/ui")
	}
	portStr := strconv.FormatInt(int64(*port), 10)
	addr := fmt.Sprintf("%s:%s", *host, portStr)
	log.Printf("Router Listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

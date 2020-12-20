package server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/markbates/pkger"
	"log"
)

func AddUiStaticServer(router *httprouter.Router, baseRoute string) {
	fs := pkger.Dir("/ui/build")
	var route string
	if baseRoute == "" || baseRoute == "/" {
		log.Print("Serving UI at /")
		route = "/*filepath"
	} else {
		log.Printf("Serving UI at %s", baseRoute)
		route = fmt.Sprintf("%s/*filepath", baseRoute)
	}
	router.ServeFiles(route, fs)
}

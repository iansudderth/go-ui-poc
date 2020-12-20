package server

import (
	"fmt"
	"github.com/elazarl/goproxy"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func UIProxy(port int, host string) http.Handler {
	proxy := goproxy.NewProxyHttpServer()
	portStr := strconv.FormatInt(int64(port), 10)
	hostStr := fmt.Sprintf("%s:%s", host, portStr)
	proxy.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Scheme = "http"
		r.URL.Host = hostStr
		proxy.ServeHTTP(w, r)
	})
	return proxy
}

func StartUIDevServer(baseRoute string, port int) {
	env := os.Environ()
	npmPath, err := exec.LookPath("npm")
	if err != nil {
		log.Fatal(err)
	}

	pathParts := []string{".", "ui"}
	uiDir, err := filepath.Abs(strings.Join(pathParts, string(os.PathSeparator)))
	if err != nil {
		log.Fatal(err)
	}

	env = append(env, PortEnv(port), SockPathEnv(baseRoute), "BROWSER=none")
	npm := exec.Cmd{
		Path: npmPath,
		Args: []string{"npm", "start"},
		Env:  env,
		Dir:  uiDir,
	}
	npm.Stdout = os.Stdout
	npm.Stderr = os.Stderr
	log.Print("Starting npm")
	err = npm.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func AddUIProxyRoutes(router *httprouter.Router, baseRoute string, proxyHost string, proxyPort int) {
	proxy := UIProxy(proxyPort, proxyHost)

	if baseRoute != "" {
		router.Handler(http.MethodGet, baseRoute, proxy)
		router.Handler(http.MethodConnect, baseRoute, proxy)
	}

	var route string
	if baseRoute == "/" {
		route = "/*uiRoute"
	} else {
		route = baseRoute + "/*uiRoute"
	}
	router.Handler(http.MethodGet, route, proxy)
	router.Handler(http.MethodConnect, route, proxy)
}

func PortEnv(port int) string {
	portString := strconv.FormatInt(int64(port), 10)
	envString := fmt.Sprintf("PORT=%s", portString)
	return envString
}

func SockPathEnv(baseRoute string) string {
	var sockPath string
	if baseRoute == "/" {
		sockPath = "/sockjs-node"
	} else {
		sockPath = fmt.Sprintf("%s/sockjs-node", baseRoute)
	}
	sockPathEnv := fmt.Sprintf("WDS_SOCKET_PATH=%s", sockPath)
	return sockPathEnv
}

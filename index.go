package main

import (
	"flag"
	"github.com/curt-labs/ariesjeep/controllers"
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"os"
)

var (
	listenAddr = flag.String("http", ":3000", "http listen address")
)

var m *martini.Martini

func main() {
	flag.Parse()

	m = martini.New()

	os.Setenv("PORT", *listenAddr)

	// Setup Middleware
	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	// m.Use(gzip.All())
	m.Use(martini.Static("static"))

	r := martini.NewRouter()

	r.Get("/(.*)", controllers.Index)

	m.Action(r.Handle)

	log.Printf("Starting server on 127.0.0.1:%s\n", *listenAddr)
	log.Println(http.ListenAndServe(*listenAddr, m))
}

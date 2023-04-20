package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	port = flag.Int("port", 4444, "Listen port")
	addr = flag.String("addr", "[::]", "IP to listen on")
	q2   = flag.String("q2", "127.0.0.1:27910", "The ip:port of the GTV server")
)

func main() {
	flag.Parse()
	RunHTTPServer()
}

func RunHTTPServer() {
	address := fmt.Sprintf("%s:%d", *addr, *port)
	router := LoadRoutes()
	httpsrv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening for web requests on", address)
	log.Println("  GTV server address:", *q2)
	log.Fatal(httpsrv.ListenAndServe())
}

func LoadRoutes() *mux.Router {
	r := mux.NewRouter()
	//r.PathPrefix(routes.Static).Handler(http.FileServer(http.Dir("./website")))
	r.HandleFunc("/", HandlerIndex)
	return r
}

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing...")
}

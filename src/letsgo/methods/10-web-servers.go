/**
Web servers
Package http serves HTTP requests using any value that implements http.Handler:

package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
In this example, the type Hello implements http.Handler.

Visit http://localhost:4000/ to see the greeting.
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

type Bottle struct{}

func (bottle Bottle) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	vals := make(map[string]string)
	vals["\nHost"] = req.Host
	vals["\nRemoteAddr"] = req.RemoteAddr
	vals["\nRequestUri"] = req.RequestURI+"\n"
	fmt.Fprint(resp, vals)
}

func main() {
	bottle := Bottle{}
	err := http.ListenAndServe("localhost:4000", bottle)
	if err != nil {
		log.Fatal(err)
	}
}

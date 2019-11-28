package main

import (
	"log"
	"os"

	"github.com/valyala/gorpc"
)

func main() {
	port := "12345"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	s := &gorpc.Server{
		// Accept clients on this TCP address.
		Addr: ":" + port,

		// Echo handler - just return back the message
		// we received from the client
		Handler: func(clientAddr string, request interface{}) interface{} {
			log.Printf("Obtained request %+v from the client %s\n",
				request, clientAddr)
			return request
		},
	}
	if err := s.Serve(); err != nil {
		log.Fatalf("Cannot start rpc server: %s", err)
	}
}

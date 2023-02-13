package api

import (
	"net"
	"net/http"
	"time"
)

func Start(db DBInterface) error {
	s := newServer(db)
	handler := corsSettings().Handler(s.router)
	s.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	//serv.logger.Info("Server is start!")
	return s.httpServer.Serve(listener)
}

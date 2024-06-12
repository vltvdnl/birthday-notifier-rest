package httpserver

import "net/http"

type Server struct {
	server *http.Server
	notify chan error
}

func New(handler http.Handler, address string) *Server {
	httpserver := &http.Server{
		Handler: handler, Addr: address,
	}

	s := &Server{
		server: httpserver,
		notify: make(chan error, 1),
	}
	s.start()
	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

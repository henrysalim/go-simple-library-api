package server

import (
	"net/http"
	"simple-library-api/internal/handlers"
)

type Server struct {
	addr        string
	bookHandler *handlers.BookHandler
}

func NewServer(addr string, bookHandler *handlers.BookHandler) *Server {
	return &Server{
		addr:        addr,
		bookHandler: bookHandler,
	}
}

func (s *Server) Run() error {
	mux := http.NewServeMux()

	//	register routes
	mux.HandleFunc("GET /books", s.bookHandler.GetBooks)
	mux.HandleFunc("POST /books/create", s.bookHandler.CreateBook)
	mux.HandleFunc("PUT /books/{id}/update", s.bookHandler.UpdateBook)
	mux.HandleFunc("DELETE /books/{id}/delete", s.bookHandler.DeleteBook)

	server := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	return server.ListenAndServe()
}

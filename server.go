package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	*http.Server
	Config ServerConfig
}

type ServerConfig struct {
	WLClient *Client
}

func NewServer(conf ServerConfig) *Server {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)

	server := &Server{
		Server: &http.Server{
			Addr:         ":8080",
			Handler:      r,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		Config: conf,
	}

	r.Get("/", server.Homepage)
	r.Get("/search", server.Search)

	return server
}

func (s *Server) Homepage(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, new(interface{}))
}

func (s *Server) Search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	characterName := r.Form.Get("name")
	if characterName == "" {
		// TODO handle
	}

	serverName := r.Form.Get("server")
	if serverName == "" {
		// TODO handle
	}

	parses, err := s.Config.WLClient.GetParses(ParsesQuery{
		CharacterName: characterName,
		Server:        serverName,
		Region:        "US",
	})

	if err != nil {
		// TODO handle
	}

	resultsTemplate.Execute(w, parses)
}

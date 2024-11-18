package gs_server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Host *Player

	Players []*Player
	Teams   []*Team

	Round int
	Turn  int
}

func NewServer() *Server {
	return &Server{
		Host:    nil,
		Players: []*Player{},
		Teams:   []*Team{},
		Round:   0,
		Turn:    0,
	}
}

func (s *Server) Listen() error {
	http.HandleFunc("/", http.FileServer(http.Dir("web")).ServeHTTP)
	http.HandleFunc("/ws", s.handleWS)
	fmt.Println("Server listening on :8080")
	return http.ListenAndServe(":8080", nil)
}

func (s *Server) handleWS(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) NextTurn() {
	if s.Turn > len(s.Teams) {
		s.Turn = 0
		s.Round++
	} else {
		s.Turn++
	}
}

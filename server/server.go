package server

import (
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"io"
)

func NewLogger() *zerolog.Logger {
	var output io.Writer

	output = zerolog.NewConsoleWriter()
	log := zerolog.New(output).With().Timestamp().Logger()

	return &log
}

func New(log *zerolog.Logger, users usersInterface) *Server {
	s := &Server{
		api:   fasthttp.Server{Name: "api"},
		log:   log,
		users: users,
	}

	s.api.Logger = s.log
	s.api.Handler = s.newRouter()

	return s
}

func (s *Server) Start() {
	const apiAddr = ":8080"

	s.log.Info().Str("addr", apiAddr).Msg("starting api http server")

	if err := s.api.ListenAndServe(apiAddr); err != nil {
		s.log.Fatal().Err(err).Msg("can't start api server")
	}
}

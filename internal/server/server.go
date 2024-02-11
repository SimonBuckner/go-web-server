package server

import (
	"github.com/labstack/echo/v4"
	"github.com/simonbuckner/go-web-server/ui"
)

type Server struct {
	Echo   *echo.Echo
	Logger echo.Logger
	UI     *ui.UI
	Pages  *Pages
	host   string
	port   string
}

func New(host, port string) *Server {
	e := echo.New()

	ui := ui.New("go-web-server")

	RegisterMiddleWare(e)

	out := Server{
		Echo:   e,
		Logger: e.Logger,
		UI:     ui,
		Pages:  RegisterPageHandlers(ui, e.Logger),
		host:   host,
		port:   port,
	}

	RegisterRoutes(&out)

	return &out
}

func (s *Server) Start() {
	s.Logger.Fatal(s.Echo.Start(s.host + ":" + s.port))
}

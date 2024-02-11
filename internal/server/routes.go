package server

func RegisterRoutes(s *Server) {

	s.Echo.Static("/css", "./static/css")
	s.Echo.Static("/js", "./static/js")
	s.Echo.Static("/", "./static/favicon")

	s.Echo.GET("/", s.Pages.Home)

}

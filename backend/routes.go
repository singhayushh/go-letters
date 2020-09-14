package backend

func (server *Server) initRoutes() {
	server.Router.GET("/", RenderHome)
	server.Router.POST("/mail", Mail)
}

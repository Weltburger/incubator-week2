package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"simple-api/internal/controller"
)

type Server struct {
	Router *echo.Echo
	Controller *controller.Controller
}

func NewServer() *Server {
	server := &Server{
		Router: echo.New(),
		Controller: controller.New(),
	}

	server.Router.Use(middleware.Logger())
	server.Router.Use(middleware.Recover())

	server.Router.GET("/", handleHome)
	apiGroup := server.Router.Group("/api")

	apiGroup.POST("/register", server.Controller.UserController().CreateUser)
	apiGroup.POST("/login", server.Controller.UserController().LogIn)
	apiGroup.GET("get/users", server.Controller.UserController().GetUsers)

	apiGroup.POST("/create/post", server.Controller.PostController().CreatePost)
	apiGroup.GET("/get/post/:id", server.Controller.PostController().GetPost)
	apiGroup.GET("/get/posts", server.Controller.PostController().GetPosts)
	apiGroup.PUT("/update/post/:id", server.Controller.PostController().UpdatePost)
	apiGroup.DELETE("/delete/post/:id", server.Controller.PostController().DeletePost)

	apiGroup.POST("/create/comment", server.Controller.CommentController().CreateComment)
	apiGroup.GET("/get/comment/:id", server.Controller.CommentController().GetComment)
	apiGroup.PUT("/update/comment/:id", server.Controller.CommentController().UpdateComment)
	apiGroup.DELETE("/delete/comment/:id", server.Controller.CommentController().DeleteComment)

	return server
}

func handleHome(c echo.Context) error {
	return c.HTML(http.StatusOK, `<h1>The Server is running</h1>`)
}



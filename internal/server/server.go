package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"satlab-api/internal/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	port int

	db database.Service
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Renderer = Renderer()
	e.Static("/static", "web/static")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", s.healthHandler)
	e.GET("/api/articles", s.ArticleGetAllPublicHandler)
	e.GET("/api/articles/:id", s.ArticleGetOneHandler)
	e.POST("/api/articles/create", s.ArticleCreateHandler)
	e.PUT("/api/articles/update", s.ArticleUpdateHandler)
	e.DELETE("/api/articles/delete", s.ArticleDeleteHandler)

	e.GET("/admin", s.AdminHomeView)
	e.Any("/admin/login", s.AdminLoginView)
	e.Any("/admin/logout", s.AdminLogoutView)
	e.GET("/admin/edit-article", s.AdminArticleEditView)

	return e
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		db: database.New(),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

package server

import (
	"net/http"
	"os"
	"satlab-api/internal/articles"

	"github.com/labstack/echo/v4"
)

func (s *Server) AdminLoginPage(c echo.Context) error {
	if c.Request().Method == http.MethodGet {
		return c.Render(http.StatusOK, "login.html", nil)
	}
	if c.FormValue("password") != os.Getenv("ADMIN_PASSWORD") {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized",
		})
	}
	if c.FormValue("username") != os.Getenv("ADMIN_USERNAME") {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized",
		})
	}
	s.CreateSession(c, "admin")
	return c.Redirect(http.StatusSeeOther, "/admin")
}

func (s *Server) AdminLogoutPage(c echo.Context) error {
	s.DestroySession(c)
	return c.Render(http.StatusOK, "logout.html", nil)
}

func (s *Server) AdminHomePage(c echo.Context) error {
	userid, err := s.GetSession(c)
	if err != nil || userid != "admin" {
		return c.Redirect(http.StatusSeeOther, "/admin/login")
	}
	repo := articles.Repository(s.db)
	articles, err := repo.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, "home.html", articles)
}

func (s *Server) AdminArticleEditPage(c echo.Context) error {
	userid, err := s.GetSession(c)
	if err != nil || userid != "admin" {
		return c.Redirect(http.StatusSeeOther, "/admin/login")
	}
	id := c.QueryParam("id")
	repo := articles.Repository(s.db)
	article, err := repo.GetOneById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, "edit_article.html", article)
}

package server

import (
	"net/http"
	"os"
	"satlab-api/internal/articles"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func checkPasswordHash(password, hash string) bool {
	println(password)
	println(os.Getenv("ADMIN_PASSWORD_HASH"))
	pwd, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	println(string(pwd[:]))
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *Server) AdminLoginView(c echo.Context) error {
	if c.Request().Method == http.MethodGet {
		return c.Render(http.StatusOK, "login.html", nil)
	}
	authorized := checkPasswordHash(c.FormValue("password"), os.Getenv("ADMIN_PASSWORD_HASH"))
	if !authorized {
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

func (s *Server) AdminLogoutView(c echo.Context) error {
	s.DestroySession(c)
	return c.Render(http.StatusOK, "logout.html", nil)
}

func (s *Server) AdminHomeView(c echo.Context) error {
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

func (s *Server) AdminArticleEditView(c echo.Context) error {
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

/*
func (s *Server) UploadFileView(c echo.Context) error {
	title := c.FormValue("title")
	a_id := c.FormValue("article_id")
	image, err := c.FormFile("image")
	if err != nil {
		return err
	}
	repo := articles.Repository(s.db)
	repo.AddImageToArticle(a_id, title, path)
}
*/

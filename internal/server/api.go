package server

import (
	"net/http"
	"satlab-api/internal/articles"

	"github.com/labstack/echo/v4"
)

func (s *Server) ArticleGetAllHandler(c echo.Context) error {
	userid, err := s.GetSession(c)
	if err != nil || userid != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized",
		})
	}
	repo := articles.Repository(s.db)
	articles_, err := repo.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, articles_)
}

func (s *Server) ArticleGetAllPublicHandler(c echo.Context) error {
	repo := articles.Repository(s.db)
	articles_, err := repo.GetAllPublic()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, articles_)
}

func (s *Server) ArticleGetOneHandler(c echo.Context) error {
	repo := articles.Repository(s.db)
	article, err := repo.GetOneById(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, article)
}

func (s *Server) ArticleCreateHandler(c echo.Context) error {
	userid, err := s.GetSession(c)
	if err != nil || userid != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized",
		})
	}
	repo := articles.Repository(s.db)
	var article articles.Article
	if err = c.Bind(&article); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := repo.CreateOne(&article); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "created",
	})
}

func (s *Server) ArticleDeleteHandler(c echo.Context) error {
	userid, err := s.GetSession(c)
	if err != nil || userid != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized",
		})
	}
	id := c.QueryParam("id")
	repo := articles.Repository(s.db)
	if err := repo.DeleteOne(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "deleted",
	})
}

func (s *Server) ArticleUpdateHandler(c echo.Context) error {
	userid, err := s.GetSession(c)
	if err != nil || userid != "admin" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized",
		})
	}
	id := c.QueryParam("id")
	var article articles.Article
	if err = c.Bind(&article); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	repo := articles.Repository(s.db)
	if err := repo.UpdateOne(id, &article); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "updated",
	})
}

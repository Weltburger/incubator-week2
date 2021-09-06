package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-api/pkg/models"
	"strconv"
)

type CommentController struct {
	controller *Controller
}

func (commentController *CommentController) CreateComment(c echo.Context) error {
	comment := new(models.Comment)
	if err := c.Bind(comment); err != nil {
		return err
	}
	commentController.controller.DB.GormDB.Create(comment)

	return c.JSON(http.StatusCreated, comment)
}

func (commentController *CommentController) GetComment(c echo.Context) error {
	key := c.Param("id")
	comment := new(models.Comment)
	commentController.controller.DB.GormDB.First(comment, key)

	return c.JSON(http.StatusOK, comment)
}

func (commentController *CommentController) UpdateComment(c echo.Context) error {
	key := c.Param("id")
	comment := new(models.Comment)
	if err := c.Bind(comment); err != nil {
		return err
	}
	comment.ID, _ = strconv.Atoi(key)
	commentController.controller.DB.GormDB.Save(comment)

	return c.JSON(http.StatusCreated, comment)
}

func (commentController *CommentController) DeleteComment(c echo.Context) error {
	key := c.Param("id")
	id, _ := strconv.ParseInt(key, 10, 64)
	comment := new(models.Comment)
	commentController.controller.DB.GormDB.Where("id = ?", id).Delete(comment)

	return c.String(http.StatusOK, "Deleted comment with id = " + key)
}

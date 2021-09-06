package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"simple-api/pkg/models"
	"strconv"
)

type PostController struct {
	controller *Controller
}

func (postController *PostController) CreatePost(c echo.Context) error {
	post := new(models.Post)
	if err := c.Bind(post); err != nil {
		return err
	}
	postController.controller.DB.GormDB.Create(post)

	return c.JSON(http.StatusCreated, post)
}

func (postController *PostController) GetPosts(c echo.Context) error {
	var posts []*models.Post
	postController.controller.DB.GormDB.Find(&posts)
	result, err := json.Marshal(posts)
	if err != nil {
		log.Fatal(err)
	}

	return c.String(http.StatusOK, string(result))
}

func (postController *PostController) GetPost(c echo.Context) error {
	key := c.Param("id")
	post := &models.Post{}
	postController.controller.DB.GormDB.First(post, key)

	return c.JSON(http.StatusOK, post)
}

func (postController *PostController) UpdatePost(c echo.Context) error {
	key := c.Param("id")
	post := new(models.Post)
	if err := c.Bind(post); err != nil {
		return err
	}
	post.ID, _ = strconv.Atoi(key)
	postController.controller.DB.GormDB.Save(post)

	return c.JSON(http.StatusCreated, post)
}

func (postController *PostController) DeletePost(c echo.Context) error {
	key := c.Param("id")
	id, _ := strconv.ParseInt(key, 10, 64)
	post := new(models.Post)
	postController.controller.DB.GormDB.Where("id = ?", id).Delete(post)

	return c.String(http.StatusOK, "Deleted post with id = " + key)
}

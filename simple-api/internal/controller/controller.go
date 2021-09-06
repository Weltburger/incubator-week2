package controller

import (
	"simple-api/internal/store"
)

type Controller struct {
	DB *store.Database
	postController *PostController
	commentController *CommentController
	userController *UserController
}

func (controller *Controller) PostController() *PostController {
	if controller.postController != nil {
		return controller.postController
	}

	controller.postController = &PostController{controller: controller}

	return controller.postController
}

func (controller *Controller) CommentController() *CommentController {
	if controller.commentController != nil {
		return controller.commentController
	}

	controller.commentController = &CommentController{controller: controller}

	return controller.commentController
}

func (controller *Controller) UserController() *UserController {
	if controller.userController != nil {
		return controller.userController
	}

	controller.userController = &UserController{controller: controller}

	return controller.userController
}

func New() *Controller {
	db := store.GetDB()
	db.Migrate()
	return &Controller{DB: db}
}

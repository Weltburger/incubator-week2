package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
	"simple-api/pkg/models"
	"strconv"
)

type UserController struct {
	controller *Controller
}

func (userController *UserController) CreateUser(c echo.Context) error {
	user := new(models.User)
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, user)

	checkUser := new(models.User)
	userController.controller.DB.GormDB.Where("email = ?", user.Email).First(checkUser)

	if checkUser.ID != "" {
		return c.String(http.StatusOK, "this user already exist!")
	}

	user.ID, _ = userController.GetActualID()
	userController.controller.DB.GormDB.Create(user)

	return c.String(http.StatusCreated, "user has been created")
}

func (userController *UserController) LogIn(c echo.Context) error {
	user := new(models.User)
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, user)

	checkUser := new(models.User)
	userController.controller.DB.GormDB.Where("email = ?", user.Email).First(checkUser)

	if checkUser.ID == "" {
		return c.String(http.StatusOK, "wrong email!")
	}

	if checkUser.Password != user.Password {
		return c.String(http.StatusOK, "wrong password!")
	}

	return c.String(http.StatusCreated, "success log in")
}

func (userController *UserController) GetActualID() (string, error) {
	var count int64
	userController.controller.DB.GormDB.Table("users").Select("count(distinct(id))").Count(&count)
 	count++

	return strconv.FormatInt(count, 10), nil
}

func (userController *UserController) GetUsers(c echo.Context) error {
	var users []*models.User
	userController.controller.DB.GormDB.Find(&users)
	result, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	return c.String(http.StatusOK, string(result))
}

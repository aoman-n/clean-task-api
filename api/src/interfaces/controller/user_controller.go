package controller

import (
	"fmt"
	"task-api/src/interfaces"
	"task-api/src/interfaces/gateway"
	"task-api/src/usecase"
	"task-api/src/usecase/interactor"
	"task-api/src/utils/auth"
)

type UserController struct {
	UserInteractor interactor.UserInteractor
}

func NewUserController(sqlhandler interfaces.SQLHandler, validator usecase.Validator) *UserController {
	userRepository := gateway.NewUserRepository(sqlhandler)
	userInteractor := interactor.NewUserInteractor(userRepository, validator)

	return &UserController{
		UserInteractor: userInteractor,
	}
}

func (uc *UserController) Index(c interfaces.Context) {
	q := c.Query("q")

	users, err := uc.UserInteractor.Search(&interactor.UserSearchInputDS{Q: q})
	if err != nil {
		fmt.Println("user search error: ", err)
		c.JSON(500, "Internal Server Error", nil)
		return
	}

	c.JSON(200, "ok", users)
}

func (uc *UserController) Show(c interfaces.Context) {
	jwtToken := auth.GetTokenFromHeader(c.Header("Authorization"))
	userID, err := auth.DecodeJWT(jwtToken)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(500, err.Error(), nil)
		return
	}

	c.JSON(200, "ok", map[string]interface{}{"userId": userID})
}

type ErrorRes struct {
	Message string `json:"message"`
}

type OkRes struct {
	Id    int64  `json:"id"`
	Token string `json:"token"`
}

func (uc *UserController) Singup(c interfaces.Context) {
	var data interactor.UserStoreInputDS
	if err := c.Bind(&data); err != nil {
		fmt.Println("failed to decode. err: ", err)
		c.JSON(400, err.Error(), nil)
		return
	}

	id, err := uc.UserInteractor.Store(data)
	if err != nil {
		c.JSON(400, err.Error(), nil)
		return
	}

	token, _ := auth.NewJWT(id)

	c.JSON(200, "ok", OkRes{Id: id, Token: token})
}

func (uc *UserController) Login(c interfaces.Context) {
	var req interactor.UserLoginInputDS
	if err := c.Bind(&req); err != nil {
		fmt.Println("failed to decode. err: ", err)
		c.JSON(400, "bad request", nil)
		return
	}

	userID, err := uc.UserInteractor.FindByLoginNameAndVerifyPassword(req)
	if err != nil {
		fmt.Println("in controller, user: ", userID)
		c.JSON(500, "bad request", err.Error())
		return
	}

	token, err := auth.NewJWT(userID)
	if err != nil {
		fmt.Println("in controller, user: ", userID)
		c.JSON(500, err.Error(), nil)
		return
	}

	// Cookieにjwtをセット
	c.SetCookie("jwt", token)

	c.JSON(200, "ok", OkRes{Id: userID, Token: token})
}

func (uc *UserController) CookieSample(c interfaces.Context) {
	jwt, err := c.GetCookie("jwt")
	if err != nil {
		c.JSON(500, "error", nil)
		return
	}

	// TODO: jwtを返却しないように
	c.JSON(200, "ok", map[string]string{"jwt": jwt})
}

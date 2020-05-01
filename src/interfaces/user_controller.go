package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-api/src/usecase"
	"task-api/src/utils/auth"
)

type UserController struct {
	UserInteractor usecase.UserInteractor
}

func NewUserController(sqlhandler SQLHandler, validator usecase.Validator) *UserController {
	userRepository := NewUserRepository(sqlhandler)
	userInteractor := usecase.NewUserInteractor(userRepository, validator)

	return &UserController{
		UserInteractor: userInteractor,
	}
}

func (uc *UserController) Index(c Context) {
	q := c.Query("q")

	users, err := uc.UserInteractor.Search(&usecase.UserSearchInputDS{Q: q})
	if err != nil {
		fmt.Println("user search error: ", err)
		c.JSON(500, "Internal Server Error", nil)
		return
	}

	c.JSON(200, "ok", users)
}

func (uc *UserController) Show(c Context) {
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

func (uc *UserController) Singup(c Context) {
	var data usecase.UserStoreInputDS
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

// Params: loginName, password
// Return: id, jwtToken
func (uc *UserController) Login(c Context) {
	var input usecase.UserLoginInputDS
	if err := c.Bind(&input); err != nil {
		fmt.Println("failed to decode. err: ", err)
		c.JSON(400, "bad request", nil)
		return
	}

	userID, err := uc.UserInteractor.FindByLoginNameAndVerifyPassword(input)
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

	c.JSON(200, "ok", OkRes{Id: userID, Token: token})
}

func jsonView(w http.ResponseWriter, code int, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf=8")
	w.WriteHeader(code)
	w.Write(b)
	return nil
}

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

func (uc *UserController) Index(w http.ResponseWriter, _ *http.Request, ps Params) {
	w.WriteHeader(200)
	w.Write([]byte("hello, world!!!"))
}

func (uc *UserController) Show(w http.ResponseWriter, r *http.Request, ps Params) {
	id := ps.ByName("id")
	fmt.Println("id: ", id)
	jwtToken := auth.GetTokenFromHeader(r)
	userID, err := auth.DecodeJWT(jwtToken)
	if err != nil {
		fmt.Println("err: ", err)
		jsonView(w, 500, "error")
		return
	}

	jsonView(w, 200, userID)
}

type ErrorRes struct {
	Message string `json:"message"`
}

type OkRes struct {
	Id    int64  `json:"id"`
	Token string `json:"token"`
}

func (uc *UserController) Singup(w http.ResponseWriter, r *http.Request, ps Params) {
	var data usecase.UserStoreInputDS
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println("failed to decode. err: ", err)
		jsonView(w, 400, ErrorRes{"error"})
		return
	}

	id, err := uc.UserInteractor.Store(data)
	if err != nil {
		jsonView(w, 400, ErrorRes{"error"})
		return
	}

	token, _ := auth.NewJWT(id)

	jsonView(w, 200, OkRes{Id: id, Token: token})
}

// Params: loginName, password
// Return: id, jwtToken
func (uc *UserController) Login(w http.ResponseWriter, r *http.Request, ps Params) {
	var input usecase.UserLoginInputDS
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		fmt.Println("failed to decode. err: ", err)
		jsonView(w, 400, ErrorRes{"error"})
		return
	}
	fmt.Println("in controller, input: ", input)

	userID, err := uc.UserInteractor.FindByLoginNameAndVerifyPassword(input)
	if err != nil {
		fmt.Println("in controller, user: ", userID)
		jsonView(w, 400, ErrorRes{"error"})
		return
	}
	fmt.Println("user: ", userID)

	token, err := auth.NewJWT(userID)
	if err != nil {
		fmt.Println("in controller, user: ", userID)
		jsonView(w, 500, ErrorRes{"error"})
		return
	}

	jsonView(w, 200, OkRes{Id: userID, Token: token})
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

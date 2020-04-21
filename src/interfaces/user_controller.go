package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-api/src/usecase"
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

func (uc *UserController) Show(w http.ResponseWriter, _ *http.Request, ps Params) {
	w.WriteHeader(200)
	w.Write([]byte("hello, world!!!"))
}

type ErrorRes struct {
	Message string `json:"message"`
}

type OkRes struct {
	Id int64 `json:"id"`
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request, ps Params) {
	var data usecase.UserStoreInputDS
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println("failed to decode. err: ", err)
		jsonView(w, 400, ErrorRes{"error"})
	}

	id, err := uc.UserInteractor.Store(data)
	if err != nil {
		jsonView(w, 400, ErrorRes{"error"})
		return
	}

	jsonView(w, 200, OkRes{Id: id})
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

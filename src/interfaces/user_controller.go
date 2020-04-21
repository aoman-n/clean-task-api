package interfaces

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	SqlHandler SQLHandler
}

func NewUserController(sqlhandler SQLHandler) *UserController {
	return &UserController{
		SqlHandler: sqlhandler,
	}
}

func (uc *UserController) Index(w http.ResponseWriter, _ *http.Request, ps Params) {
	w.WriteHeader(200)
	w.Write([]byte("hello, world!!!"))
}

func (uc *UserController) Show(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("hello, world!!!"))
}

func (uc *UserController) Create(w http.ResponseWriter, _ *http.Request, ps Params) {
	w.WriteHeader(200)
	w.Write([]byte("hello, world!!!"))
}

package interfaces

import (
	"net/http"
	"task-api/src/infrastructure/db"

	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	SqlHandler *db.Sqlhandler
}

func NewUserController(sqlhandler *db.Sqlhandler) *UserController {
	return &UserController{
		SqlHandler: sqlhandler,
	}
}

func (uc *UserController) Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("hello, world!!!"))
}

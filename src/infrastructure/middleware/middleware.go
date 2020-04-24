package middleware

import (
	"net/http"
	"task-api/src/interfaces"
	"task-api/src/usecase"
	"task-api/src/utils/auth"
)

type Middlewares struct {
	userRepository usecase.UserRepository
}

func New(sqlhandler interfaces.SQLHandler) *Middlewares {
	userRepository := interfaces.NewUserRepository(sqlhandler)

	return &Middlewares{
		userRepository: userRepository,
	}
}

func (m *Middlewares) Authenticate(h interfaces.HttpHandlerWithUserID) interfaces.HttpHandler {
	return func(w http.ResponseWriter, r *http.Request, params interfaces.Params) {
		token := auth.GetTokenFromHeader(r)
		userId, err := auth.DecodeJWT(token)
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("unauthorized."))
			return
		}

		h(w, r, params, userId)
	}
}

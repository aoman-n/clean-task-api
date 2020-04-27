package middleware

import (
	"net/http"
	"strconv"
	"task-api/src/entity/model"
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

// JWTの検証とuserIDのinject
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

// projectに参加しているか確認するmiddleware
func (m *Middlewares) RequiredJoinedProject(h interfaces.HttpHandlerWithUserID) interfaces.HttpHandlerWithUserID {
	return func(w http.ResponseWriter, r *http.Request, params interfaces.Params, uID int64) {
		projectID, _ := strconv.Atoi(params.ByName("id"))
		usersInProject, err := m.userRepository.FindByProjectID(projectID)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("internal server error."))
			return
		}

		// プロジェクトに参加しているユーザーからのリクエストの場合コントローラの処理を行う
		for _, u := range *usersInProject {
			if u.ID == uID {
				h(w, r, params, uID)
				return
			}
		}

		w.WriteHeader(401)
		w.Write([]byte("project not joined. cannot create task."))
		return
	}
}

// projectに参加+書き込み権限があるか確認するmiddleware
func (m *Middlewares) RequiredWriteRole(h interfaces.HttpHandlerWithUserID) interfaces.HttpHandlerWithUserID {
	return func(w http.ResponseWriter, r *http.Request, params interfaces.Params, uID int64) {
		projectID, _ := strconv.Atoi(params.ByName("id"))
		usersInProject, err := m.userRepository.FindByProjectID(projectID)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("internal server error."))
			return
		}

		// プロジェクトに参加してるかつ、Write以上の権限がある場合にControllerの処理を行う
		for _, u := range *usersInProject {
			if u.ID == uID && (u.Role == model.Admin || u.Role == model.Write) {
				h(w, r, params, uID)
				return
			}
		}

		w.WriteHeader(401)
		w.Write([]byte("Project does not have permission"))
		return
	}
}

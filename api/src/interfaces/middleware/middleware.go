package middleware

import (
	"net/http"
	"strconv"
	"task-api/src/interfaces"
	"task-api/src/interfaces/gateway"
	"task-api/src/usecase/interactor"
	"task-api/src/utils/auth"
)

type Middleware interface {
	Auth(c interfaces.Context)
	RequiredJoinedProject(c interfaces.Context)
	RequiredWriteRole(c interfaces.Context)
}

type middleware struct {
	midInteractor interactor.MiddlewareInteractor
}

func NewMiddlewre(sqlhandler interfaces.SQLHandler) Middleware {
	userRepo := gateway.NewUserRepository(sqlhandler)
	midInteractor := interactor.NewMiddlewareInteractor(userRepo)

	return &middleware{midInteractor}
}

func (m *middleware) Auth(c interfaces.Context) {
	token := auth.GetTokenFromHeader(c.Header("Authorization"))
	userId, err := auth.DecodeJWT(token)
	if err != nil {
		c.JSON(401, "unauthorized.", nil)
		c.Abort()
		return
	}

	c.Set("userId", userId)
}

func (m *middleware) RequiredJoinedProject(c interfaces.Context) {
	uID := c.MustGet("userId").(int64)
	pID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		// TODO: error handling
	}

	canAccess, err := m.midInteractor.CanAccessProject(&interactor.CanAccessProjectInputDS{
		UID: uID,
		PID: pID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error.", nil)
		c.Abort()
		return
	}
	if !canAccess {
		c.JSON(http.StatusUnauthorized, "project not joined. cannot create task.", nil)
		c.Abort()
		return
	}

	return
}

func (m *middleware) RequiredWriteRole(c interfaces.Context) {
	uID := c.MustGet("userId").(int64)
	pID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		// TODO: error handling
	}

	canWrite, err := m.midInteractor.CanWriteProject(&interactor.CanWriteProjectInputDS{
		UID: uID,
		PID: pID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error.", nil)
		c.Abort()
		return
	}
	if !canWrite {
		c.JSON(http.StatusUnauthorized, "project does not have permission", nil)
		c.Abort()
		return
	}

	return
}

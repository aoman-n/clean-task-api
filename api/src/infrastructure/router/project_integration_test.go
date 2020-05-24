// +build integration

package router

import (
	"net/http/httptest"
	"strings"
	"task-api/src/infrastructure/db"
	"task-api/src/infrastructure/engine"
	"task-api/src/interfaces"
	"task-api/src/interfaces/controller"
	"task-api/src/interfaces/middleware"
	"task-api/src/testutil"
	"task-api/src/utils/auth"
	"testing"

	"gopkg.in/go-playground/validator.v9"
)

func setUpTx(t *testing.T) (interfaces.SQLHandler, func()) {
	t.Helper()

	sqlhandler := db.NewSqlhandler()
	tx, err := sqlhandler.Conn.Begin()
	if err != nil {
		t.Fatal("failed to begin transaction err: ", err)
	}
	txHandler := &db.SQLHandler{Tx: tx}

	return txHandler, func() {
		tx.Rollback()
	}
}

func setUpProjectController(tx interfaces.SQLHandler) (*middleware.Middleware, *controller.ProjectController) {
	validator := validator.New()
	mi := middleware.NewMiddlewre(tx)
	co := controller.NewProjectController(tx, validator)

	return mi, co
}

func TestProject_IndexIntegration(t *testing.T) {

	// txHandlerのセットアップ
	tx, tearDown := setUpTx(t)
	defer tearDown()

	// seedデータ投入
	testutil.ExecSchema(tx, "./testdata/project_router/index_seed.sql")

	// jwt作成
	jwt, err := auth.NewJWT(1)
	if err != nil {
		t.Fatal("failed to create jwt err: ", err)
	}

	// router設定
	mi, co := setUpProjectController(tx)
	router := engine.New()
	router.GET("/projects", mi.Auth, co.Index)
	req := httptest.NewRequest("GET", "/projects", nil)
	req.Header.Add("Authorization", "Bearer "+jwt)
	w := httptest.NewRecorder()

	// リクエスト実行
	router.ServeHTTP(w, req)

	t.Log(strings.Repeat("=", 50))
	t.Log("status: ", w.Code)
	t.Log("body: ", w.Body.String())
	t.Log("header content type: ", w.Header().Get("Content-Type"))
	t.Log(strings.Repeat("=", 50))

	testutil.AssertResponse(t, w, 200, "./testdata/project_router/index_ok_response.golden")
}

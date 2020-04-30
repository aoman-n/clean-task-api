package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"task-api/src/infrastructure/middleware"
	"task-api/src/interfaces"
	"task-api/src/usecase"

	"github.com/julienschmidt/httprouter"
)

func log(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[ACCESS] ", r.Method, r.URL, r.Host, r.RequestURI)
}

func logging(h interfaces.HttpHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Println("[ACCESS] ", r.Method, r.URL, r.Host, r.RequestURI)
		h(w, r, params)
	}
}

type Ctx interface {
	JSON(code int, msg string, data interface{})
	Param(key string) string
	Set(key string, value interface{})
	Get(key string) (value interface{}, exists bool)
	MustGet(key string) (value interface{})
}

type ctx struct {
	w  http.ResponseWriter
	r  *http.Request
	ps httprouter.Params
}

type Response struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func (c *ctx) JSON(code int, msg string, data interface{}) {
	res := &Response{
		Msg:  msg,
		Data: data,
	}

	b, err := json.Marshal(res)
	if err != nil {
		c.error()
		return
	}
	c.w.Header().Set("Content-Type", "application/json; charset=utf=8")
	c.w.WriteHeader(code)
	c.w.Write(b)
}

func (c *ctx) error() {
	c.w.WriteHeader(http.StatusInternalServerError)
	c.w.Write([]byte("Internal server error."))
}

func (c *ctx) Param(key string) string {
	return c.ps.ByName(key)
}

func (c *ctx) Set(key string, value interface{}) {
	c.r.Context()
	ctx := context.WithValue(c.r.Context(), key, value)
	c.r = c.r.WithContext(ctx)
}

func (c *ctx) Get(key string) (value interface{}, exists bool) {
	v := c.r.Context().Value(key)
	return v, v != nil
}

func (c *ctx) MustGet(key string) interface{} {
	if value, exists := c.Get(key); exists {
		return value
	}

	panic("Key \"" + key + "\" does not exist")
}

type Handle interface {
	Dispatch(handlers []HandlerFunc) httprouter.Handle
}

type handle struct{}

func (h *handle) Dispatch(handlers []HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log(w, r)
		ctx := ctx{w, r, ps}
		for _, h := range handlers {
			h(&ctx)
		}
	}
}

type Engine struct {
	router *httprouter.Router
	handle Handle
}

func NewHRouter() *Engine {
	router := httprouter.New()
	handle := &handle{}

	return &Engine{router, handle}
}

type HandlerFunc func(c Ctx)

func (e *Engine) GET(path string, handlers ...HandlerFunc) {
	e.router.GET(path, e.handle.Dispatch(handlers))
}

func (e *Engine) POST(path string, handlers ...HandlerFunc) {
	e.router.POST(path, e.handle.Dispatch(handlers))
}

func (e *Engine) PUT(path string, handlers ...HandlerFunc) {
	e.router.PUT(path, e.handle.Dispatch(handlers))
}

func (e *Engine) DELETE(path string, handlers ...HandlerFunc) {
	e.router.GET(path, e.handle.Dispatch(handlers))
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.router.ServeHTTP(w, r)
}

func injectUserID(c Ctx) {
	// rand.Seed(time.Now().Unix())
	// number := rand.Intn(100)
	// fmt.Println("################## number: ", number)

	// if number%2 == 0 {
	// c.Set("userID", 100)
	// }
}

func hello(c Ctx) {
	// uID := c.Get("userID").(int)
	// uID := c.Get("userID")
	uID := c.MustGet("userID")
	fmt.Println("uID: ", uID)

	c.JSON(http.StatusOK, "ok", "hello, world")
}

func Handler(sqlhandler interfaces.SQLHandler, validator usecase.Validator) *Engine {
	hrouter := NewHRouter()

	hrouter.GET("/hello", injectUserID, hello)

	middlewares := middleware.New(sqlhandler)

	userController := interfaces.NewUserController(sqlhandler, validator)
	projectController := interfaces.NewProjectController(sqlhandler, validator)
	taskController := interfaces.NewTaskController(sqlhandler, validator)
	tagController := interfaces.NewTagController(sqlhandler, validator)

	router := httprouter.New()
	/* users API */
	router.POST("/signup", logging(userController.Singup))
	router.POST("/login", logging(userController.Login))
	router.GET("/users", logging(userController.Index))
	router.GET("/users/:id", logging(userController.Show))

	/* projects API */
	router.GET("/projects", logging(middlewares.Authenticate(projectController.Index)))
	router.POST("/projects", logging(middlewares.Authenticate(projectController.Create)))
	router.DELETE("/projects/:id", logging(middlewares.Authenticate(projectController.Delete)))

	/* task API */
	router.GET("/projects/:id", logging(middlewares.Authenticate(middlewares.RequiredJoinedProject(taskController.Index))))
	router.POST("/projects/:id/tasks", logging(middlewares.Authenticate(middlewares.RequiredWriteRole(taskController.Create))))
	router.DELETE("/projects/:id/tasks/:task_id", logging(middlewares.Authenticate(middlewares.RequiredWriteRole(taskController.Delete))))
	router.PUT("/projects/:id/tasks/:task_id", logging(middlewares.Authenticate(middlewares.RequiredWriteRole(taskController.Update))))

	/* tags API */
	router.GET("/projects/:id/tags", logging(middlewares.Authenticate(middlewares.RequiredWriteRole(tagController.Index))))
	router.POST("/projects/:id/tags", logging(middlewares.Authenticate(middlewares.RequiredWriteRole(tagController.Create))))
	router.PUT("/tags/:id", logging(middlewares.Authenticate(tagController.Update)))
	router.DELETE("/tags/:id", logging(middlewares.Authenticate(tagController.Delete)))

	return hrouter
}

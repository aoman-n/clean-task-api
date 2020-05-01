package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// type Context interface {
// 	JSON(code int, msg string, data interface{})
// 	Param(key string) string
// 	Set(key string, value interface{})
// 	Get(key string) (value interface{}, exists bool)
// 	MustGet(key string) (value interface{})
// }

type Context struct {
	w      http.ResponseWriter
	r      *http.Request
	ps     httprouter.Params
	enable bool
}

type Response struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func (c *Context) JSON(code int, msg string, data interface{}) {
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

func (c *Context) error() {
	c.w.WriteHeader(http.StatusInternalServerError)
	c.w.Write([]byte("Internal server error."))
}

func (c *Context) Param(key string) string {
	return c.ps.ByName(key)
}

func (c *Context) Set(key string, value interface{}) {
	c.r.Context()
	Context := context.WithValue(c.r.Context(), key, value)
	c.r = c.r.WithContext(Context)
}

func (c *Context) Get(key string) (value interface{}, exists bool) {
	v := c.r.Context().Value(key)
	return v, v != nil
}

func (c *Context) MustGet(key string) interface{} {
	if value, exists := c.Get(key); exists {
		return value
	}

	panic("Key \"" + key + "\" does not exist")
}

func (c *Context) Bind(value interface{}) error {
	// TODO: validtorを使ってvalidationする
	return json.NewDecoder(c.r.Body).Decode(value)
}

func (c *Context) Query(key string) string {
	return c.r.URL.Query().Get("q")
}

func (c *Context) Header(key string) string {
	return c.r.Header.Get(key)
}

func (c *Context) Abort() {
	c.enable = false
}

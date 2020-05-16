package interfaces

type Context interface {
	JSON(code int, msg string, data interface{})
	Param(key string) string
	Set(key string, value interface{})
	Get(key string) (value interface{}, exists bool)
	MustGet(key string) (value interface{})
	Bind(value interface{}) error
	Query(key string) string
	Header(key string) string
	Abort()
	SetCookie(name, value string)
	GetCookie(name string) (string, error)
}

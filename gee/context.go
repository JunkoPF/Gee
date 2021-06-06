package gee

import "net/http"

type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{}
}

func (c *Context) PostForm(key string) string {

}

func (c *Context) Query(key string) string {

}

func (c *Context) Status(code int) {

}

func (c *Context) SetHeader(key string, value string) {

}

func (c *Context) String(code int, format string, values ...interface{}) {

}

func (c *Context) JSON(code int, obj interface{}) {

}

func (c *Context) Data(code int, data []byte) {

}

func (c *Context) HTML(code int, html string) {

}

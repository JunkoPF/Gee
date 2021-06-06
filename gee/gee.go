package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {

}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
}

func (engine *Engine) Run(addr string) (err error) {
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {


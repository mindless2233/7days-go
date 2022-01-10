package cin

import "net/http"

type Engine struct {
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (e *Engine) Run(port string) {
	http.ListenAndServe(port, e)
}

package cin

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(context *Context) {

		t := time.Now()
		context.Next()
		log.Printf("sys-log:[%d] %s  %v", context.StatusCode, context.Req.RequestURI, time.Since(t))
	}

}

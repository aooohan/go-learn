package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func trace(msg string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])
	var str strings.Builder
	str.WriteString(msg + "\n Traceback:")
	// 获取调用栈信息
	frames := runtime.CallersFrames(pcs[:n])
	for {
		next, more := frames.Next()
		str.WriteString(fmt.Sprintf("\n\t%s:%d", next.File, next.Line))
		if !more {
			break
		}
	}

	//for _, pc := range pcs[:n] {
	//	fn := runtime.FuncForPC(pc)
	//	file, line := fn.FileLine(pc)
	//	str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	//}
	return str.String()

}
func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if e := recover(); e != nil {
				msg := fmt.Sprintf("%s", e)
				log.Printf("%s\n\n", trace(msg))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}

package gee

import (
	"log"
	"net/http"
	"path"
	"strings"
)

// HandlerFunc defines the req handler used by gee
type HandlerFunc func(*Context)

type (
	// Engine implement the interface of ServeHTTP
	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup // store all groups
	}

	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc // support middleware
		engine      *Engine       // all groups share a Engine instance
	}
)

func (r *RouterGroup) Group(prefix string) *RouterGroup {
	engine := r.engine
	newGroup := &RouterGroup{
		prefix: r.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (r *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(r.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.StatusCode = http.StatusOK
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

func (r *RouterGroup) Static(relativePath string, root string) {
	handler := r.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	r.Get(urlPattern, handler)
}

func (r *RouterGroup) Use(handlers ...HandlerFunc) {
	r.middlewares = append(r.middlewares, handlers...)
}

func (r *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := r.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	r.engine.router.addRoute(method, pattern, handler)
}

// New create a new Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func Default() *Engine {
	engine := New()
	engine.Use(Recovery(), Logger())
	return engine
}

func (r *RouterGroup) Get(pattern string, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

func (r *RouterGroup) Post(pattern string, handler HandlerFunc) {
	r.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := newContext(w, r)
	middlewares := make([]HandlerFunc, 0)
	for _, group := range e.groups {
		if strings.HasPrefix(context.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	context.handlers = middlewares
	e.router.handle(context)
}

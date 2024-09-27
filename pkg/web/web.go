package web

import "github.com/gin-gonic/gin"

type Route interface {
	Register(*gin.Engine)
}

type Server struct {
	*gin.Engine
	Router []Route // 路由组
}

// New 初始化服务器
func New(e *gin.Engine) *Server {
	return &Server{
		Engine: e,
	}
}

// RegisterRouter 注册路由组
func (s *Server) RegisterRouter(rs ...Route) {
	s.Router = append(s.Router, rs...)
	for _, r := range rs {
		r.Register(s.Engine)
	}
}

// Limit 设置全局并发限制
func (s *Server) Limit(max int) {
	s.Engine.Use(LimitMiddleware(max))
}

// LimitMiddleware 并发限制中间件
func LimitMiddleware(max int) gin.HandlerFunc {
	sem := make(chan struct{}, max)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }

	return func(ctx *gin.Context) {
		acquire()       // before request
		defer release() // after request
		ctx.Next()
	}
}

package web_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/miajio/aide/pkg/web"
)

type userRouter struct{}

var UserRouter web.Route = (*userRouter)(nil)

func (u *userRouter) Register(e *gin.Engine) {
	e.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func TestWeb(t *testing.T) {
	w := web.New(gin.Default())
	w.Limit(64)

	w.RegisterRouter(UserRouter)
	w.Run(":8080")
}

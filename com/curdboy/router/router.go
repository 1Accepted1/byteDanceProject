package router

import (
	"byteDanceProject/com/curdboy/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {



	g := r.Group("/tiktok")

	//session 提示注册的信息无法识别
	//gob.Register(types.Teacher)
	//gob.Register(types.Admin)
	//gob.Register(types.Student)

	//g.Use(sessions.Sessions(types.SessionName,types.Store))

	// 成员管理
	//g.POST("/member/create", controller.MemberCreatePost)
	//g.GET("/member", controller.MemberGetOne)
	//g.GET("/member/list",controller.MemberGetList)
	//g.POST("/member/update",controller.MemberUpdate)
	//g.POST("/member/delete",controller.MemberDelete)

	//g.Handle("GET", "/say_hello", func(c *gin.Context) {
	//	name := c.Query("name")
	//	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	//})
	g.GET("/hello",controller.Test)


}

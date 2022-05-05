package main

import (
	"byteDanceProject/com/curdboy/router"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {

	//启动redis
	//myRedis.RedisService = redis.NewClient(&redis.Options{Network: myRedis.RedisNetWork, Addr: myRedis.RedisAddr + ":" + myRedis.RedisPort, IdleTimeout: myRedis.RedisTimeOutForKeep})

	//定义一个路由
	r := gin.Default()

	//导入模板
	//r.LoadHTMLGlob("view/*")

	//导入静态资源
	//r.Static("/static", "./static")

	//注册路由
	router.RegisterRouter(r)
	//连接字节的数据库
	//utils.InitMysql()

	//可以取消表名的复数形式，使得表名和结构体名称一致
	//utils.Db.SingularTable(true)
	r.Handle("GET", "/say_hello", func(c *gin.Context) {
		name := c.Query("name")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	r.Run(":80")

}

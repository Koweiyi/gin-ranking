package routes

import (
	"go-ranking/controllers"
	// logger "go-ranking/pkg"

	"go-ranking/config"
	logger "go-ranking/pkg"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"

	sessions_redis "github.com/gin-contrib/sessions/redis"
)

func GetRoute() *gin.Engine {
	r := gin.Default()

	// 引用日志工具
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	// redis 配置
	store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAdress, "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	user := r.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)
		// user.GET("/info/:id/:name", controllers.UserController{}.GetUserInfo)
		// user.POST("/list", controllers.UserController{}.GetList)
		// user.GET("/err", controllers.UserController{}.MakeError)

		// user.PUT("/add", func(ctx *gin.Context) {
		// 	ctx.String(200, "user add")
		// })

		// user.DELETE("/delete", func(ctx *gin.Context) {
		// 	ctx.String(200, "user delete")
		// })
	}

	return r
}

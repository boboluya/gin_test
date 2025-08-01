package main

import (
	"gin_simulate_auth/user/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	{
		userPath := g.Group("/user")
		userPath.POST("/add", handler.Add)
		userPath.POST("/update", handler.Update)
		userPath.POST("/delete", handler.Delete)
		userPath.GET("/selectList", handler.GetList)
		userPath.GET("/selectOne", handler.GetOne)
		userPath.GET("/selectById", handler.GetOneById)
	}
	err := g.Run(":8080")
	if err != nil {
		return
	}
}

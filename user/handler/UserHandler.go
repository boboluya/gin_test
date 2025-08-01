package handler

import (
	"gin_simulate_auth/user/data"
	"gin_simulate_auth/user/mapper"
	"gin_simulate_auth/user/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var userMapper = mapper.UserMapperFromMap{}

var userService = service.NewUserService(userMapper)

func Add(c *gin.Context) {
	var user data.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	add := userService.Add(user)
	log.Println("add:", add)
	c.JSON(200, gin.H{
		"msg":  "添加成功",
		"data": add,
	})

}

func Update(c *gin.Context) {
	var user data.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	update, err := userService.Update(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "更新成功",
		"data": update,
	})
}

func Delete(c *gin.Context) {
	var user data.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	deleted := userService.Delete(user)
	c.JSON(200, gin.H{
		"msg":  "删除成功",
		"data": deleted,
	})
}

func GetList(c *gin.Context) {
	var user data.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list := userService.SelectList(user)
	c.JSON(200, gin.H{
		"msg":  "200",
		"data": list,
	})
}

func GetOne(c *gin.Context) {
	var user data.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	row, err := userService.SelectOne(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data": row,
	})
}

func GetOneById(c *gin.Context) {
	id := c.GetInt64("id")
	byId, err := userService.SelectOneById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data": byId,
	})
}

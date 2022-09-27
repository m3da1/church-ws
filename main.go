package main

import (
	"church-ws/controller"
	"church-ws/models"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	if err := models.Connect(); err != nil {
		fmt.Printf("database error: %+v", err.Error())
		panic(err)
	}
	r.GET("churchdbws/users", controller.FindUsers)
	r.GET("churchdbws/user/:id", controller.FindUser)
	r.POST("churchdbws/user", controller.CreateUser)
	r.PUT("churchdbws/user", controller.UpdateUser)
	r.DELETE("churchdbws/user/:id", controller.DeleteUser)
	r.POST("churchdbws/user/login", controller.Login)
	r.POST("churchdbws/user/changepasswd", controller.ChangePassword)

	r.GET("churchdbws/members", controller.FindMembers)
	r.GET("churchdbws/member/:id", controller.FindMember)
	r.POST("churchdbws/member", controller.CreateMember)
	r.PUT("churchdbws/member", controller.UpdateMember)
	r.DELETE("churchdbws/member/:id", controller.DeleteMember)

	r.GET("churchdbws/stewardgroups", controller.FindStewardGroups)
	r.GET("churchdbws/stewardgroup/:id", controller.FindStewardGroup)
	r.POST("churchdbws/stewardgroup", controller.CreateStewardGroup)
	r.PUT("churchdbws/stewardgroup", controller.UpdateStewardGroup)
	r.DELETE("churchdbws/stewardgroup/:id", controller.DeleteStewardGroup)

	r.Run()
}

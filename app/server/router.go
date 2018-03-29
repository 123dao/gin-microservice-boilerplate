package server

import (
	//"fmt"

	"github.com/burxtx/gin-microservice-boilerplate/app/config"
	"github.com/burxtx/gin-microservice-boilerplate/app/endpoints"
	"github.com/burxtx/gin-microservice-boilerplate/app/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter(eps endpoints.AppEndpoint) *gin.Engine {
	c := config.GetConfig()
	if c.GetString("env.mode") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(endpoints.HealthEndpoint)
	router.GET("/health", health.Status)
	auth := new(endpoints.AuthEndpoint)
	router.GET("/logout", auth.Logout)

	router.Use(middlewares.CasAuthMiddleware())
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/index.html")
	})
	router.GET("/user", auth.GetCurrentUser)
	v1 := router.Group("v1")
	{
		auditGroup := v1.Group("app")
		{
			auditGroup.GET("/get", eps.GetEndpoint)

		}
	}
	return router
}

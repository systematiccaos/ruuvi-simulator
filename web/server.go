package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", pingHandler)
		cfggrp := v1.Group("conf")
		{
			cfggrp.GET("list", listConfHandler)
		}
		accdatagrp := v1.Group("acc-data")
		{
			accdatagrp.GET("data/:tag", getAccDataHandler)
			accdatagrp.GET("list", listAccTagsHandler)
		}
		structuregrp := v1.Group("structure")
		{
			structuregrp.GET("list", listNodesHandler)
		}
	}
	// r.GET("/ping", pingHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.InstanceName("swagger")))

}

//	@BasePath	/api/v1

// PingExample godoc
//
//	@Summary	ping example
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	pong
//	@Router			/ping [get]
func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

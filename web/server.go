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
		cfggrp := v1.Group("config")
		{
			cfggrp.GET("get/:gatewayid", getConfHandler)
			cfggrp.GET("get/:gatewayid/:tagaddress", getTagConfHandler)
		}
		accdatagrp := v1.Group("acc-data")
		{
			accdatagrp.GET("get/:tag", getLatestDataTagsHandler)
			accdatagrp.GET("get/:tag/:page", getAccDataHandler)
		}
		structuregrp := v1.Group("structure")
		{
			structuregrp.GET("gateway/list", listNodesHandler)
			structuregrp.GET("tag/list/:gateway", listTagsHandler)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.InstanceName("swagger")))

}

//	@BasePath	/api/v1

// PingExample godoc
//
//	@Summary	ping example
//	@Schemes
//	@Description	do ping
//	@Tags			ping
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	pong
//	@Router			/ping [get]
func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

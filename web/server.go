package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/systematiccaos/ruuvi-simulator/mock"
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

//	@BasePath	/api/v1

// List godoc
//
//	@Summary	lists available configs
//	@Schemes
//	@Description	get configs of gateways to be able to decide on version updates etc.
//	@Tags			config
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	json-conf
//	@Router			/list [get]
func listConfHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

//	@BasePath	/api/v1

// List godoc
//
//	@Summary	lists available gateways
//	@Schemes
//	@Description	lists gateways
//	@Tags			config
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	json-conf
//	@Router			/list [get]
func listNodesHandler(c *gin.Context) {
	gws := mock.GetMock().Gateways
	c.JSON(http.StatusOK, gws)
}

//	@BasePath	/api/v1/acc-data

// List godoc
//
//	@Summary	gets acc data of the specified tag
//	@Schemes
//	@Description	gets data of the specified tag - get your tags via "list" first
//	@Tags			config
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	json-conf
//	@Param			tag	path		string	true	"the tags name"	example("ruuvi_1234")
//	@Router			/data/{tag} [get]
func getAccDataHandler(c *gin.Context) {
	tag := c.Param("tag")

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"tag":     tag,
	})
}

//	@BasePath	/api/v1/acc-data

// List godoc
//
//	@Summary	gets acc data of the specified tag
//	@Schemes
//	@Description	gets data of the specified tag - get your tags via "list" first
//	@Tags			config
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	json-conf
//	@Router			/list [get]
func listAccTagsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

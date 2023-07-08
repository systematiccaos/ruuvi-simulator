package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@BasePath	/api/v1

// List godoc
//
//	@Summary	gets acc data of the specified tag
//	@Schemes
//	@Description	gets data of the specified tag - get your tags via "list" first
//	@Tags			acc-data
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	json-conf
//	@Param			tag	path		string	true	"the tags name"	example("ruuvi_1234")
//	@Router			/acc-data/{tag} [get]
func getAccDataHandler(c *gin.Context) {
	tag := c.Param("tag")

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"tag":     tag,
	})
}

//	@BasePath	/api/v1

// List godoc
//
//	@Summary	gets acc data of the specified tag
//	@Schemes
//	@Description	gets data of the specified tag - get your tags via "list" first
//	@Tags			acc-data
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.AccelerationSensor
//	@Router			/acc-data/list [get]
func listAccTagsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

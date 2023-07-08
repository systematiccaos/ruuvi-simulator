package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@BasePath	/api/v1

// List godoc
//
//	@Summary	lists available configs
//	@Schemes
//	@Description	get configs of gateways to be able to decide on version updates etc.
//	@Tags			config
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.GatewayConfig
//	@Router			/config/list [get]
func listConfHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/systematiccaos/ruuvi-simulator/mock"
	"github.com/systematiccaos/ruuvi-simulator/model"
)

//	@BasePath	/api/v1

// List godoc
//
//	@Summary	lists available configs
//	@Schemes
//	@Description	get config of a specific gateway to be able to decide on version updates etc. To find the gateways id use /structure/gateway/list
//	@Tags			config
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.GatewayConfig
//	@Param	gatewayid	path	string	true	"the gateways id"
//	@Router			/config/get/{gatewayid} [get]
func getConfHandler(c *gin.Context) {
	var gwc model.GatewayConfig
	mck := mock.GetMock()
	id := c.Param("gatewayid")
	for i := range mck.Gateways {
		if mck.Gateways[i].ID == id {
			gwc = mck.Gateways[i].Config
			c.JSON(http.StatusOK, gin.H{
				"config": gwc,
			})
			return
		}
	}
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": "no content for this id",
	})
}

//	@BasePath	/api/v1

// List godoc
//
//	@Summary	lists available configs
//	@Schemes
//	@Description	get config of a specific gateway to be able to decide on version updates etc. To find the gateways id use /structure/gateway/list, to get the tags address use /structure/tag/list/{gatewayid}
//	@Tags			config
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.GatewayConfig
//	@Param	gatewayid	path	string	true	"the gateways id"
//	@Param	tagaddress	path	string	true	"the tags id"
//	@Router			/config/get/{gatewayid}/{tagaddress} [get]
func getTagConfHandler(c *gin.Context) {
	var tc model.TagConfig
	mck := mock.GetMock()
	id := c.Param("gatewayid")
	tid := c.Param("tagaddress")
	for i := range mck.Gateways {
		if mck.Gateways[i].ID == id {
			for idx := range mck.Gateways[i].Tags {
				if mck.Gateways[i].Tags[idx].Address == tid {
					tc = mck.Gateways[i].Tags[idx].Config
					c.JSON(http.StatusOK, gin.H{
						"config": tc,
					})
					return
				}
			}
		}
	}
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": "no content for this id",
	})
}

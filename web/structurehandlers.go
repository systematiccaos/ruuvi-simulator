package web

import (
	"net/http"
	"reflect"

	"github.com/barkimedes/go-deepcopy"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/systematiccaos/ruuvi-simulator/mock"
	"github.com/systematiccaos/ruuvi-simulator/model"
)

//	@BasePath	/api/v1

// List godoc
//
//	@Summary	lists available gateways
//	@Schemes
//	@Description	lists gateways
//	@Tags			structure
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.Gateway
//	@Router			/structure/gateway/list [get]
func listNodesHandler(c *gin.Context) {
	mck := mock.GetMock()
	clone := mock.Mock{}
	cloneif, err := deepcopy.Anything(mck.Gateways)
	if err != nil {
		logrus.Errorln(err)
	}
	clone.Gateways = reflect.ValueOf(cloneif).Interface().([]model.Gateway)
	for i := range clone.Gateways {
		clone.Gateways[i].Tags = nil
	}
	// for i := range clone.Gateways {
	// 	clone.Gateways[i].LastContact = mck.Gateways[i].LastContact
	// 	for idx := range clone.Gateways[i].Tags {
	// 		clone.Gateways[i].Tags[idx].Sensors = []model.Sensor{}
	// 		clone.Gateways[i].Tags[idx].LastContact = mck.Gateways[i].Tags[idx].LastContact
	// 		clone.Gateways[i].Tags[idx].Online = mck.Gateways[i].Tags[idx].Online
	// 	}
	// }
	c.JSON(http.StatusOK, clone)
}

//	@BasePath	/api/v1

// List godoc
//
//	@Summary	lists available tags
//	@Schemes
//	@Description	lists tags
//	@Tags			structure
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.Gateway
//	@Router			/structure/tag/list/{gateway_id} [get]
func listTagsHandler(c *gin.Context) {
	gwid := c.Param("gateway")
	mck := mock.GetMock()
	clone := mock.Mock{}
	cloneif, err := deepcopy.Anything(mck.Gateways)
	if err != nil {
		logrus.Errorln(err)
	}
	gw := model.Gateway{}
	clone.Gateways = reflect.ValueOf(cloneif).Interface().([]model.Gateway)
	for i := range clone.Gateways {
		clone.Gateways[i].LastContact = mck.Gateways[i].LastContact
		for idx := range clone.Gateways[i].Tags {
			clone.Gateways[i].Tags[idx].Sensors = []model.Sensor{}
			clone.Gateways[i].Tags[idx].LastContact = mck.Gateways[i].Tags[idx].LastContact
			clone.Gateways[i].Tags[idx].Online = mck.Gateways[i].Tags[idx].Online
		}
		if clone.Gateways[i].ID == gwid {
			gw = clone.Gateways[i]
		}
	}
	c.JSON(http.StatusOK, gw.Tags)
}

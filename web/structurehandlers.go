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
//	@Router			/structure/list [get]
func listNodesHandler(c *gin.Context) {
	mockcl := mock.GetMock()
	cloneif, err := deepcopy.Anything(mockcl)
	if err != nil {
		logrus.Errorln(err)
	}
	clone := reflect.ValueOf(cloneif).Elem().Interface().(mock.Mock)
	for i := range clone.Gateways {
		for idx := range clone.Gateways[i].Tags {
			clone.Gateways[i].Tags[idx].Sensors = []model.Sensor{&model.AccelerationSensor{Tp: "AccelerationSensor"}}
		}
	}
	c.JSON(http.StatusOK, clone.Gateways)
}

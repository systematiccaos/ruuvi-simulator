package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/systematiccaos/ruuvi-simulator/mock"
	"github.com/systematiccaos/ruuvi-simulator/model"
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
//	@Param			tag	path		string	true	"the tags address"	example("12:34:56:78:90:12")
//	@Router			/acc-data/get/{tag}/{page} [get]
func getAccDataHandler(c *gin.Context) {
	address := c.Param("tag")
	pagei := 1
	page := c.Param("page")
	sizei := 10
	size := c.Query("size")

	if size != "" {
		var err error
		sizei, err = strconv.Atoi(size)
		if err != nil {
			logrus.Errorln(err)
		}
	}
	if page != "" {
		var err error
		pagei, err = strconv.Atoi(page)
		if err != nil {
			logrus.Errorln(err)
		}
	}

	mck := mock.GetMock()
	var tag model.Tag
	found := false
	for _, t := range mck.Tags {
		if t.Address == address {
			tag = t
			found = true
		}
	}
	if !found {
		c.JSON(500, gin.H{"error": "could not find a tag with the given address"})
		return
	}

	measurements := []model.Measurement{}
	for _, s := range tag.Sensors {
		measurements = append(measurements, s.GetMeasurements()...)
	}
	nextpage := pagei + 1
	if len(measurements) <= (pagei-1)*sizei+sizei {
		lp_offset := int((len(measurements) - 1) / 10)
		lp_size := 10
		pagei = lp_offset + 1
		sizei = lp_size
		logrus.Println(pagei, sizei)
		nextpage = 1
	}

	c.JSON(http.StatusOK, gin.H{
		"address":      address,
		"page":         pagei,
		"size":         sizei,
		"next_page":    nextpage,
		"measurements": measurements[(pagei-1)*sizei : (pagei-1)*sizei+sizei],
	})
}

//	@BasePath	/api/v1

// List godoc
//
//	@Summary	gets latest data of the specified tag
//	@Schemes
//	@Description	gets the latest data of the specified tag - get your tags via "list" first
//	@Tags			acc-data
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.AccelerationSensor
//	@Param			tag	path		string	true	"the tags address"	example("12:34:56:78:90:12")
//	@Router			/acc-data/get/{tag} [get]
func getLatestDataTagsHandler(c *gin.Context) {
	address := c.Param("tag")
	pagei := 1
	page := c.Param("page")
	sizei := 10
	size := c.Query("size")

	if size != "" {
		var err error
		sizei, err = strconv.Atoi(size)
		if err != nil {
			logrus.Errorln(err)
		}
	}
	if page != "" {
		var err error
		pagei, err = strconv.Atoi(page)
		if err != nil {
			logrus.Errorln(err)
		}
	}

	mck := mock.GetMock()
	var tag model.Tag
	found := false
	for _, t := range mck.Tags {
		if t.Address == address {
			tag = t
			found = true
		}
	}
	if !found {
		c.JSON(500, gin.H{"error": "could not find a tag with the given address"})
		return
	}

	var latest_measurements []model.Measurement
	for _, s := range tag.Sensors {
		latest_measurements = append(latest_measurements, s.GetMeasurements()[len(s.GetMeasurements())-1])
	}
	c.JSON(http.StatusOK, gin.H{
		"address":      address,
		"page":         pagei,
		"size":         sizei,
		"measurements": latest_measurements,
	})
}

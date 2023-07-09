package main

import (
	"math/rand"

	_ "github.com/systematiccaos/ruuvi-simulator/docs"
	"github.com/systematiccaos/ruuvi-simulator/mock"

	"github.com/gin-gonic/gin"
	"github.com/systematiccaos/going-forward/util"
	"github.com/systematiccaos/ruuvi-simulator/web"
)

func main() {
	util.SetupLogs()
	// waitch := make(chan bool)
	// go calcPendulums()
	rand.Seed(1010)
	r := gin.Default()
	web.SetupRoutes(r)
	m := mock.GetMock()
	go m.Run()
	r.Run()
	// <-waitch
}

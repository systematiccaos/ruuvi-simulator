package main

import (
	"math/rand"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
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
	seedno := 1010
	if os.Getenv("SEED") != "" {
		var err error
		seedno, err = strconv.Atoi(os.Getenv("SEED"))
		if err != nil {
			logrus.Fatalln(err)
		}
	}
	rand.Seed(int64(seedno))
	r := gin.Default()
	web.SetupRoutes(r)
	m := mock.GetMock()
	go m.Run()
	r.Run()
	// <-waitch
}

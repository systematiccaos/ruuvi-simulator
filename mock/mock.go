package mock

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/systematiccaos/ruuvi-simulator/model"
)

var instance *Mock

type Mock struct {
	Gateways []model.Gateway `json:"gateways,omitempty"`
	Tags     []model.Tag     `json:"tags,omitempty"`
}

func GetMock() *Mock {
	if instance != nil {
		return instance
	}
	instance = startMock()
	return instance
}

func startMock() *Mock {
	mock := Mock{}
	for i := 0; i < rand.Intn(100); i++ {
		gw := model.NewGateway()
		mock.Gateways = append(mock.Gateways, gw)
		mock.Tags = append(mock.Tags, gw.Tags...)
	}
	logrus.Println(mock)
	return &mock
}

func (m *Mock) Run() {
	i := 0
	for {
		i++
		for _, gw := range m.Gateways {
			gw.Update()
		}
		time.Sleep(time.Second / 60)
	}
}

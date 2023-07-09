package mock

import (
	"math/rand"
	"time"

	"github.com/systematiccaos/ruuvi-simulator/model"
	"github.com/systematiccaos/ruuvi-simulator/utils"
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
	for i := 0; i < rand.Intn(50); i++ {
		gw := model.NewGateway()
		mock.Gateways = append(mock.Gateways, gw)
		mock.Tags = append(mock.Tags, gw.Tags...)
	}
	// logrus.Println(mock)
	return &mock
}

func (m *Mock) Run() {
	i := 0
	for {
		i++
		for i := range m.Gateways {
			m.Gateways[i].Update()
			for idx := range m.Gateways[i].Tags {
				if m.Gateways[i].Tags[idx].WantsChange {
					newidx := i + 1
					if i+1 < len(m.Gateways[i].Tags) {
						newidx = 0
					}
					m.Gateways[newidx].Tags = append(m.Gateways[i].Tags, m.Gateways[i].Tags[idx])
					utils.DeleteItem[model.Tag](m.Gateways[i].Tags, idx)
				}
			}
		}
		time.Sleep(time.Second / 60)
	}
}

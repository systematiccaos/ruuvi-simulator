package model

import "math/rand"

type Gateway struct {
	Tags   []Tag         `json:"tags"`
	Config GatewayConfig `json:"gateway_config"`
}

func NewGateway() Gateway {
	gw := Gateway{}
	gw.Tags = []Tag{}
	tagcnt := rand.Intn(50)
	for i := 0; i < tagcnt; i++ {
		tag := NewTag()
		gw.Tags = append(gw.Tags, tag)
	}
	return gw
}

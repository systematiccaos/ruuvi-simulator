package model

import (
	"math/rand"
	"time"
)

type Gateway struct {
	Tags        []Tag         `json:"tags"`
	Config      GatewayConfig `json:"gateway_config"`
	LastContact time.Time     `json:"last_contact"`
	Online      bool          `json:"online"`
}

func NewGateway() Gateway {
	gw := Gateway{}
	gw.Online = !(rand.Float32() < 0.2)
	gw.LastContact = time.Now()
	gw.Tags = []Tag{}
	tagcnt := rand.Intn(5)
	for i := 0; i < tagcnt; i++ {
		tag := NewTag()
		gw.Tags = append(gw.Tags, tag)
	}
	return gw
}

func (g *Gateway) Update() {
	for _, tag := range g.Tags {
		tag.Update()
	}
	if g.Online {
		g.LastContact = time.Now()
	}
	if !(rand.Float32() < 0.05) {
		g.Online = !g.Online
	}
}

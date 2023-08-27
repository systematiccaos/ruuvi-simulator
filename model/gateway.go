package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

// Gateway - stores information on the gateway runtime
// @Description tags - all the tags, known to the gateway
// @Description config - GatewayConfig that belongs to this Gateway
// @Description network_segment - subnet the Gateway is in
// @Description last_contact - last time the api heard back from the Gateway
// @Description online - bool that determines if the Gateway is currently online
// @Description ip_address - current IPv4 of the Gateway
// @Description id - unique identifier
type Gateway struct {
	Tags           []Tag         `json:"tags,omitempty"`
	Config         GatewayConfig `json:"-"`
	NetworkSegment int           `json:"network_segment"`
	LastContact    time.Time     `json:"last_contact"`
	Online         bool          `json:"online"`
	IPAdress       string        `json:"ip_address"`
	ID             string        `json:"id"`
}

func NewGateway() Gateway {
	gw := Gateway{}
	gw.Online = !(rand.Float32() < 0.2)
	gw.LastContact = time.Now()
	gw.Tags = []Tag{}
	tagcnt := rand.Intn(15)
	gw.NetworkSegment = rand.Intn(5)
	gw.IPAdress = fmt.Sprintf("%d.%d.%d.%d", 10, 0, gw.NetworkSegment, rand.Intn(254))
	hash := md5.New().Sum([]byte(gw.IPAdress))
	gw.ID = hex.EncodeToString(hash[:])
	for i := 0; i < tagcnt; i++ {
		tag := NewTag()
		gw.Tags = append(gw.Tags, tag)
	}
	return gw
}

func (g *Gateway) Update() {
	if g.Online {
		g.LastContact = time.Now()
	}
	if rand.Float32() < 0.00005 {
		g.Online = !g.Online
		g.NetworkSegment = rand.Intn(5)
	}
	if !g.Online {
		return
	}
	for i := range g.Tags {
		g.Tags[i].Update()
	}
}

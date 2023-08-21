package model

import (
	"fmt"
	"math/rand"
)

// GatewayConfig - stores information on the gateway
// @Description version - version of the gateway's firmware
// @Description poll_interval - how often the tags will be polled via bluetooth
// @Description max_allowed_clients - how many tags will be accepted
// @Description api_timeout - how long it takes for the gateway to restart when the api-server is unavailable
type GatewayConfig struct {
	Version           string  `json:"version"`
	PollInterval      int     `json:"poll_interval"`
	MaxAllowedClients int     `json:"max_allowed_clients"`
	APITimeout        float64 `json:"api_timeout"`
}

func NewGatewayConfig() GatewayConfig {
	poll_interval := 10
	max_clients := 100
	api_timeout := 10000.0
	if rand.Float32() < 0.1 {
		poll_interval = 100
	}
	if rand.Float32() < 0.1 {
		max_clients = 50
	}
	if rand.Float32() < 0.1 {
		api_timeout = 1000
	}

	cfg := GatewayConfig{
		Version:           fmt.Sprintf("1.1.%d", rand.Intn(3)),
		PollInterval:      poll_interval,
		MaxAllowedClients: max_clients,
		APITimeout:        api_timeout,
	}
	return cfg
}

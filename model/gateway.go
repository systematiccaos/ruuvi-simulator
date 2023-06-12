package model

type Gateway struct {
	Tags   []Tag         `json:"tags"`
	Config GatewayConfig `json:"gateway_config"`
}

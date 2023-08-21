package model

import (
	"math/rand"
)

// Tag - stores information on the tag
// @Description samplerate - samplerate of the tag
// @Description scan_interval - interval the sensors of the tag will be polled (ms)
// @Description resolution - bit depth resolution of the sensors
// @Description scale - scaling factor for values from the sensors (for compression)
// @Description dsp_function - dsp function for signal evaluation (enum)
// @Description dsp_parameter - dsp configuration parameter (enum)
// @Description mode - current measurement mode (enum)
// @Description divider - divider for the samplerate
type TagConfig struct {
	SampleRate   int     `json:"samplerate"`
	Interval     float64 `json:"scan_interval"`
	Resolution   int     `json:"resolution"`
	Scale        int     `json:"scale"`
	DSPFunction  int     `json:"dsp_function"`
	DSPParameter int     `json:"dsp_parameter"`
	Mode         int     `json:"mode"`
	Divider      int     `json:"divider"`
}

func NewTagConfig() TagConfig {
	samplerate := 10
	poll_interval := float64(1.0 / samplerate)
	scale := 1
	if rand.Float32() < 0.1 {
		poll_interval = 100.0
	}
	if rand.Float32() < 0.1 {
		poll_interval = float64(1.0/samplerate + 10)
	}
	if rand.Float32() < 0.1 {
		scale = 1000
	}

	cfg := TagConfig{
		SampleRate:   samplerate,
		Interval:     poll_interval,
		Resolution:   rand.Intn(100),
		Scale:        scale,
		DSPFunction:  rand.Intn(10),
		DSPParameter: rand.Intn(10),
	}
	return cfg
}

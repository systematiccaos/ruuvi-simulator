package model

import (
	"fmt"
	"math/rand"
	"time"
)

// Tag - stores information on the tag
// @Description sensors - all the sensors, mounted on this tag (list)
// @Description address - unique MAC-address of the tag (bluetooth MAC)
// @Description name - name of the tag that derives from the MAC-address
// @Description last_contact - last time the Gateway heard back from the Tag
// @Description online - bool that determines if the Tag is currently online
// @Description config - TagConfig that belongs to this Tag
type Tag struct {
	Sensors     []Sensor  `json:"sensors"`
	Address     string    `json:"address"`
	Name        string    `json:"name"`
	LastContact time.Time `json:"last_contact"`
	Online      bool      `json:"online"`
	WantsChange bool      `json:"-"`
	Config      TagConfig `json:"config"`
}

func NewTag() Tag {
	sfx1 := rand.Intn(255)
	sfx2 := rand.Intn(255)
	tag := Tag{
		Name:        "ruuvi_" + fmt.Sprintf("%x%x", sfx1, sfx2),
		Address:     fmt.Sprintf("%x:%x:%x:%x:%x:%x", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255), sfx1, sfx2),
		LastContact: time.Now(),
		Online:      true,
	}
	as := NewAccelerationSensor()
	ts := NewTemperatureSensor()
	hs := NewHumiditySensor()
	vs := NewVoltageSensor()
	tag.Config = NewTagConfig()
	tag.Sensors = append(tag.Sensors, &as, &ts, &hs, &vs)
	return tag
}

func (t *Tag) GetAllSensorsMeasurements() []Measurement {
	measurements := []Measurement{}
	for _, sensor := range t.Sensors {
		measurement := sensor.GetMeasurements()
		measurements = append(measurements, measurement...)
	}
	return measurements
}

func (t *Tag) Update() {
	if t.Online {
		t.LastContact = time.Now()
		// logrus.Println(t.LastContact)
	}
	if rand.Float32() < 0.00005 {
		t.Online = !t.Online
	}
	t.WantsChange = rand.Float32() < 0.00005
	if !t.Online {
		return
	}
	for i := range t.Sensors {
		t.Sensors[i].Update()
	}
}

package config

type Telemetry struct {
	ServiceName string
}

func (c *Telemetry) SetDefaults() {
	c.ServiceName = "store-api"

}

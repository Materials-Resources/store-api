package config

type Server struct {
	Host string `koanf:"host"`
	Port int    `koanf:"port"`
}

func (c *Server) SetDefaults() {
	c.Port = 8080
}

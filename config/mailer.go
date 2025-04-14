package config

type Mailer struct {
	Host      string `koanf:"host"`
	Port      int    `koanf:"port"`
	Username  string `koanf:"username"`
	Password  string `koanf:"password"`
	Sender    string `koanf:"sender"`
	Recipient string `koanf:"recipient"`
}

func (c *Mailer) SetDefaults() {
	c.Port = 25
}

package pkg

type Autolaod struct {
	Config *Config
}

func NewAutoload(config Config) *Autolaod {
	return &Autolaod{
		Config: &config,
	}
}

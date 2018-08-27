package bizconfig

// db biz config
type DbConfig struct {
	Dialect string `yaml:"dialect"`
	Address string `yaml:"address"`
}

// session biz config
type SessionConfig struct {
}

// web biz config
type WebConfig struct {
}

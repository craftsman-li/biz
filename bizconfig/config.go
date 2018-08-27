package bizconfig

// db biz config
type DbConfig struct {
	Dialect string `yaml:"dialect"`
	Address string `yaml:"address"`
	Prefix string `yaml:"prefix"`
}

// session biz config
type SessionConfig struct {
}

type LogInfo struct {
	Level string
}

// web biz config
type WebConfig struct {
}

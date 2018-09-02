package bizconfig

// db biz config
type DbConfig struct {
	Dialect string `yaml:"dialect"`
	Address string `yaml:"address"`
	Prefix  string `yaml:"prefix"`
}

func (db *DbConfig) InitParam(defaultPrefix string) {
	if "" == db.Prefix {
		db.Prefix = defaultPrefix
	}
	if "" == db.Dialect {
		db.Dialect = "mysql"
	}
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

package bizconfig

// db biz config
type DbConfig struct {
	Dialect string `yaml:"dialect"`
	Address string `yaml:"address"`
	Prefix  string `yaml:"prefix"`
}

// session biz config
type SessionConfig struct {
}

type LogInfo struct {
	Level string
}

// web biz config
type WebConfig struct {
	// 是否校验允不允许登录
	// 如果为true, 则在账号密码验证通过后，检验是否有此路径登录的权限
	LoginCheckEnable bool
}

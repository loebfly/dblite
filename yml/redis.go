package yml

type Redis struct {
	Host     string    `yaml:"host"`
	Port     int       `yaml:"port"`
	Password string    `yaml:"password"`
	Database int       `yaml:"database"`
	Timeout  int       `yaml:"timeout"`
	Pool     RedisPool `yaml:"pool"`
}

type RedisPool struct {
	Min     int `yaml:"min"`
	Max     int `yaml:"max"`
	Idle    int `yaml:"idle"`
	Timeout int `yaml:"timeout"`
}

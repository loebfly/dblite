package yml

type Redis struct {
	Host     string    `yml:"host"`
	Port     int       `yml:"port"`
	Password string    `yml:"password"`
	Database int       `yml:"database"`
	Timeout  int       `yml:"timeout"`
	Pool     RedisPool `yml:"pool"`
}

type RedisPool struct {
	Min     int `yml:"min"`
	Max     int `yml:"max"`
	Idle    int `yml:"idle"`
	Timeout int `yml:"timeout"`
}

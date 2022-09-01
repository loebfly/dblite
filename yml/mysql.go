package yml

type Mysql struct {
	Url   string    `yaml:"url" json:"url"`
	Debug bool      `yaml:"debug" json:"debug"`
	Pool  MysqlPool `yaml:"pool" json:"pool"`
}

type MysqlPool struct {
	Max     int              `yaml:"max" json:"max"`
	Idle    int              `yaml:"idle" json:"idle"`
	Timeout MysqlPoolTimeout `yaml:"timeout" json:"timeout"`
}

type MysqlPoolTimeout struct {
	Idle int `yaml:"idle" json:"idle"`
	Life int `yaml:"life" json:"life"`
}

package yml

type Mysql struct {
	Url   string    `yml:"url" json:"url"`
	Debug bool      `yml:"debug" json:"debug"`
	Pool  MysqlPool `yml:"pool" json:"pool"`
}

type MysqlPool struct {
	Max     int              `yml:"max" json:"max"`
	Idle    int              `yml:"idle" json:"idle"`
	Timeout MysqlPoolTimeout `yml:"timeout" json:"timeout"`
}

type MysqlPoolTimeout struct {
	Idle int `yml:"idle" json:"idle"`
	Life int `yml:"life" json:"life"`
}

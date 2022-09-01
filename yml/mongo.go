package yml

type Mongo struct {
	Url      string `yaml:"url" json:"url"`
	Database string `yaml:"database" json:"database"`
	PoolMax  int    `yaml:"pool_max" json:"pool_max"`
}

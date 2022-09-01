package atest

import (
	"github.com/loebfly/dblite"
	"testing"
)

func TestRedis(t *testing.T) {
	ymlPath := "/github.com/dblite/atest/app.yml"
	//err := dblite.Init(ymlPath, dblite.UseRedis)
	err := dblite.Redis.Init(ymlPath)
	//err := dblite.Redis.InitObj(yml.Redis{
	//	Host:     "",
	//	Port:     0,
	//	Password: "",
	//	Database: 5,
	//})

	if err != nil {
		t.Error(err)
		return
	}

	rds, err := dblite.Redis.GetDB()
	if err != nil {
		t.Error(err)
		return
	}

	val := rds.Get("partner:general:id").Val()
	t.Log(val)
	dblite.Redis.SafeExit()
}

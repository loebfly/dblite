package atest

import (
	"github.com/loebfly/dblite"
	"github.com/loebfly/dblite/yml"
	"testing"
)

func TestMysql(t *testing.T) {
	//ymlPath := "/github.com/dblite/test/mysql.yml"
	//err := dblite.Init(ymlPath, dblite.UseMysql)
	//err := dblite.Mysql.Init(ymlPath)
	err := dblite.Mysql.InitObj(yml.Mysql{
		Url:   "",
		Debug: false,
		Pool: yml.MysqlPool{
			Max:  20,
			Idle: 10,
			Timeout: yml.MysqlPoolTimeout{
				Idle: 60,
				Life: 60,
			},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	db, err := dblite.Mysql.GetDB()
	if err != nil {
		t.Error(err)
		return
	}

	type IpAll struct {
		CountryCode string `gorm:"column:country_code" json:"countryCode"` // 国家代码, 部分数据可能带6
		ShortCode   string `gorm:"column:short_code" json:"shortCode"`     // 国家代码, 不带6
		StartIp     string `gorm:"column:start_ip" json:"startIp"`         // 起始IP
		Length      string `gorm:"column:length" json:"length"`            // 长度
		IsDelete    int    `gorm:"column:is_delete" json:"isDelete"`       // 状态 1 为已删除
	}
	var ip IpAll
	err = db.Table("country_ip_all").Limit(1).First(&ip).Error
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ip)
}

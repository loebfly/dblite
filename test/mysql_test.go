package test

import (
	"dblite"
	"dblite/mysql"
	"fmt"
	"testing"
)

func TestMysql(t *testing.T) {
	enter := mysql.Enter{}
	err := enter.LocalInit("/Users/luchunqing/Documents/QingGe/SourceTree/dblite/test/app.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestLite(t *testing.T) {
	err := dblite.LocalInit("/Users/luchunqing/Documents/QingGe/SourceTree/dblite/test/app.yml", dblite.UseTypeMysql)
	if err != nil {
		fmt.Println(err)
		return
	}
	db, err := dblite.Mysql.GetDB()
	if err != nil {
		return
	}

	type CountryIpAllRow struct {
		Code     string `gorm:"column:country_code" json:"countryCode"` // 国家代码, 部分数据可能带6
		StartIp  string `gorm:"column:start_ip" json:"startIp"`         // 起始IP
		Length   string `gorm:"column:length" json:"length"`            // 长度
		IsDelete int    `gorm:"column:is_delete" json:"isDelete"`       // 状态 1 为已删除
	}
	var list []CountryIpAllRow
	err = db.Table("country_ip_all").Limit(10).Find(&list).Error
	if err != nil {
		return
	}

	for _, v := range list {
		fmt.Println(v)
	}

	dblite.Mysql.SafeExit()
}

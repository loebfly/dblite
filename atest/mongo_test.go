package atest

import (
	"github.com/loebfly/dblite"
	"testing"
)

func TestMongo(t *testing.T) {
	ymlPath := "/github.com/dblite/atest/app.yml"
	err := dblite.Init(ymlPath, dblite.UseMongo)
	if err != nil {
		t.Error(err)
		return
	}

	db, closeDB, err := dblite.Mongo.GetDB()
	if err != nil {
		t.Error(err)
		return
	}

	type Country struct {
		CountryId   string `bson:"country_id"`
		CountryName string `bson:"country_name"`
	}

	var list []Country

	db.C("country").Find(nil).All(&list)
	t.Log(list)
	closeDB()
	dblite.Mongo.SafeExit()
}

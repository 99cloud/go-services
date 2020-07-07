package database

import (
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/internal/APP_46ea591951824d8e9376b0f98fe4d48a/model"
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/database"
	"reflect"
)

func init() {
	database.OrmRegisterList[reflect.TypeOf(model.GormDBTest{}).Name()] = model.GormDBTest{}
}

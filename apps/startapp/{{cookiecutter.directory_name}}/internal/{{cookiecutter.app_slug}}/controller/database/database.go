package database

import (
	"{{cookiecutter.project_slug}}/internal/{{cookiecutter.app_slug}}/model"
	"{{cookiecutter.project_slug}}/pkg/database"
	"reflect"
)

func init() {
	database.OrmRegisterList[reflect.TypeOf(model.GormDBTest{}).Name()] = model.GormDBTest{}
}

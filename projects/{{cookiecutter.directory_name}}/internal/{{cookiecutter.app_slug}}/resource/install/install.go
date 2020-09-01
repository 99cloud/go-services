package install

import (
	v1 "{{cookiecutter.project_slug}}/internal/{{cookiecutter.app_slug}}/resource/v1"
	"{{cookiecutter.project_slug}}/pkg/server/runtime"
	"github.com/emicklei/go-restful"
)

func init() {
	Install(runtime.Container)
}

func Install(c *restful.Container) {
	runtime.Must(v1.AddToContainer(c))
}

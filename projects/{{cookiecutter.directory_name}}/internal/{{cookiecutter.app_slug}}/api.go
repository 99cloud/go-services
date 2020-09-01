package {{cookiecutter.app_slug}}

import (
	resourcev1 "{{cookiecutter.project_slug}}/internal/{{cookiecutter.app_slug}}/resource/v1"
	"{{cookiecutter.project_slug}}/pkg/server/runtime"
	"github.com/emicklei/go-restful"
)

func InstallAPIs(container *restful.Container) {
	runtime.Must(resourcev1.AddToContainer(container))
}

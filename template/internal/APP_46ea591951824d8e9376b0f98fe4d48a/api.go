package APP_46ea591951824d8e9376b0f98fe4d48a

import (
	resourcev1 "PROJECT_46ea591951824d8e9376b0f98fe4d48a/internal/APP_46ea591951824d8e9376b0f98fe4d48a/resource/v1"
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/server/runtime"
	"github.com/emicklei/go-restful"
)

func InstallAPIs(container *restful.Container) {
	runtime.Must(resourcev1.AddToContainer(container))
}

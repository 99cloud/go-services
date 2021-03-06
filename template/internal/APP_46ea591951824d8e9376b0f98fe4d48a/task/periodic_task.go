package task

import (
	"fmt"

	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/cmd/APP_46ea591951824d8e9376b0f98fe4d48a/app/options"
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/logger"

	"github.com/jasonlvhit/gocron"
)

// RunPeriodicTask is to run periodic task
func RunPeriodicTask(s *options.ServerRunOptions) {
	logger.Info(nil, "Run periodic task")
	// gocron.Every(1).Second().Do(mockEcho, s)
	<-gocron.Start()
}

func getInternalEndpoint(s *options.ServerRunOptions, uri string) string {
	var endpoint string
	if s.GenericServerRunOptions.InsecurePort != 0 {
		endpoint = fmt.Sprintf("http://%v:%v%v", s.GenericServerRunOptions.BindAddress,
			s.GenericServerRunOptions.InsecurePort, uri)
	}
	if s.GenericServerRunOptions.SecurePort != 0 && len(s.GenericServerRunOptions.TlsCertFile) > 0 && len(s.GenericServerRunOptions.TlsPrivateKey) > 0 {
		endpoint = fmt.Sprintf("https://%v:%v%v", s.GenericServerRunOptions.BindAddress,
			s.GenericServerRunOptions.SecurePort, uri)
	}
	return endpoint
}

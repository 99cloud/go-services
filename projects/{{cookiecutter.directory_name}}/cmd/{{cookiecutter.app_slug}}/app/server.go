package app

import (
	"fmt"
	"net/http"

	"{{cookiecutter.project_slug}}/cmd/{{cookiecutter.app_slug}}/app/options"
	"{{cookiecutter.project_slug}}/internal/{{cookiecutter.app_slug}}"
	"{{cookiecutter.project_slug}}/internal/{{cookiecutter.app_slug}}/task"
	"{{cookiecutter.project_slug}}/pkg/client"
	"{{cookiecutter.project_slug}}/pkg/logger"
	"{{cookiecutter.project_slug}}/pkg/server"
	serverconfig "{{cookiecutter.project_slug}}/pkg/server/config"
	"{{cookiecutter.project_slug}}/pkg/server/filter"
	"{{cookiecutter.project_slug}}/pkg/server/runtime"
	"{{cookiecutter.project_slug}}/pkg/server/version"
	"{{cookiecutter.project_slug}}/pkg/utils/signals"

	"github.com/spf13/cobra"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

func NewAPIServerCommand() *cobra.Command {
	s := options.NewServerRunOptions()

	cmd := &cobra.Command{
		Use:  "{{cookiecutter.app_slug}}",
		Long: `{{cookiecutter.app_slug}} restful api server`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := serverconfig.Load()
			if err != nil {
				return err
			}

			err = Complete(s)
			if err != nil {
				return err
			}

			if errs := s.Validate(); len(errs) != 0 {
				return utilerrors.NewAggregate(errs)
			}

			go task.RunPeriodicTask(s)

			return Run(s, signals.SetupSignalHandler())
		},
	}

	fs := cmd.Flags()
	fs.StringVar(&s.Loglevel, "loglevel", s.Loglevel, "info server log level, e.g. debug,info")
	namedFlagSets := s.Flags()

	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	return cmd
}

// apply server run options to configuration
func Complete(s *options.ServerRunOptions) error {

	// loading configuration file
	conf := serverconfig.Get()

	conf.Apply(&serverconfig.Config{
		OrmOptions: s.OrmOptions,
	})

	*s = options.ServerRunOptions{
		GenericServerRunOptions: s.GenericServerRunOptions,
		OrmOptions:              conf.OrmOptions,
		Loglevel:                s.Loglevel,
	}

	return nil
}

func Run(s *options.ServerRunOptions, stopCh <-chan struct{}) error {
	logger.SetLevelByString(s.Loglevel)
	err := CreateClientSet(serverconfig.Get(), stopCh)
	if err != nil {
		return err
	}

	err = CreateAPIServer(s)
	if err != nil {
		return err
	}

	return nil
}

func CreateClientSet(conf *serverconfig.Config, stopCh <-chan struct{}) error {
	csop := &client.ClientSetOptions{}

	csop.SetOrmOptions(conf.OrmOptions)

	client.NewClientSetFactory(csop, stopCh)

	return nil
}

func CreateAPIServer(s *options.ServerRunOptions) error {
	var err error

	container := runtime.Container
	container.DoNotRecover(false)
	container.Filter(filter.Logging)
	container.RecoverHandler(server.LogStackOnRecover)

	{{cookiecutter.app_slug}}.InstallAPIs(container)

	// install config api
	serverconfig.InstallAPI(container)

	if s.GenericServerRunOptions.InsecurePort != 0 {
		logger.Info(nil, "Server [version: %s] Start on %s:%d", version.Version, s.GenericServerRunOptions.BindAddress, s.GenericServerRunOptions.InsecurePort)
		err = http.ListenAndServe(fmt.Sprintf("%s:%d", s.GenericServerRunOptions.BindAddress, s.GenericServerRunOptions.InsecurePort), container)
		if err == nil {
			logger.Info(nil, "Server listening on insecure port %d.", s.GenericServerRunOptions.InsecurePort)
		}
	}

	if s.GenericServerRunOptions.SecurePort != 0 && len(s.GenericServerRunOptions.TlsCertFile) > 0 && len(s.GenericServerRunOptions.TlsPrivateKey) > 0 {
		err = http.ListenAndServeTLS(fmt.Sprintf("%s:%d", s.GenericServerRunOptions.BindAddress, s.GenericServerRunOptions.SecurePort), s.GenericServerRunOptions.TlsCertFile, s.GenericServerRunOptions.TlsPrivateKey, container)
		if err == nil {
			logger.Info(nil, "Server listening on secure port %d.", s.GenericServerRunOptions.SecurePort)
		}
	}

	return err
}

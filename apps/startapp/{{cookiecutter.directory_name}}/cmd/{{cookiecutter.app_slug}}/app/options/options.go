package options

import (
	"{{cookiecutter.project_slug}}/pkg/client/orm"
	"{{cookiecutter.project_slug}}/pkg/client/ssh"
	genericoptions "{{cookiecutter.project_slug}}/pkg/server/options"
	cliflag "k8s.io/component-base/cli/flag"
)

type ServerRunOptions struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions

	OrmOptions *orm.OrmOptions

	SshOptions *ssh.SshOptions

	Loglevel string
}

func NewServerRunOptions() *ServerRunOptions {

	s := ServerRunOptions{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		OrmOptions:              orm.NewOrmOptions(),
		SshOptions:              ssh.NewOrmOptions(),
		Loglevel:                "info",
	}

	return &s
}

func (s *ServerRunOptions) Flags() (fss cliflag.NamedFlagSets) {

	s.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
	s.OrmOptions.AddFlags(fss.FlagSet("orm"))
	s.SshOptions.AddFlags(fss.FlagSet("ssh"))

	return fss
}

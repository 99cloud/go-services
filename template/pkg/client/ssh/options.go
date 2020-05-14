package ssh

import (
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/utils/reflectutils"

	"github.com/spf13/pflag"
)

type SshOptions struct {
	Host               string `json:"host,omitempty" yaml:"host" description:"remote server address"`
	AuthenticationMode string `json:"authentication-mode,omitempty" yaml:"authentication-mode" description:"Server authentication mode"`
	Username           string `json:"username,omitempty" yaml:"username"`
	Password           string `json:"password" yaml:"password"`
	SshPort            int    `json:"ssh-port" yaml:"ssh-port"`
	PrivateKeyFile     string `json:"private-key-file" yaml:"private-key-file"`
	PrivateKey         string `json:"private-key" yaml:"private-key"`
	AnsibleVarDir      string `json:"ansible-var-dir" yaml:"ansible-var-dir"`
	AnsibleVarFile     string `json:"ansible-var-file" yaml:"ansible-var-file"`
	AnsibleCommand     string `json:"ansible-command" yaml:"ansible-command"`
}

func NewOrmOptions() *SshOptions {
	return &SshOptions{
		Host:               "172.16.30.34",
		AuthenticationMode: "basic",
		Username:           "root",
		Password:           "password@123",
		SshPort:            22,
		PrivateKeyFile:     "/root/.ssh/id_rsa",
		PrivateKey: `-----BEGIN OPENSSH PRIVATE KEY-----
    b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAABFwAAAAdzc2gtcn
    ...
    3LRHxpjyl7IWcS8MxpXMEa9D9zcN017fdwyckLX+HjGgXiiAmoMA9Q9GIu6TqdAfuKUma9
    Wks6hF5EQzoAdn8AAAAOcm9vdEBzaW1vbnMtcGMBAgMEBQ==
    -----END OPENSSH PRIVATE KEY-----`,
		AnsibleVarDir:  "/etc/kolla/",
		AnsibleVarFile: "globals.yml",
		AnsibleCommand: "/root/testup",
	}
}

func (m *SshOptions) Validate() []error {
	var errors []error

	return errors
}

func (m *SshOptions) ApplyTo(options *SshOptions) {
	reflectutils.Override(options, m)
}

func (m *SshOptions) AddFlags(fs *pflag.FlagSet) {

	fs.StringVar(&m.Host, "ssh-host", m.Host, ""+
		"ssh remote server address")

	fs.StringVar(&m.AuthenticationMode, "ssh-auth-mode", m.AuthenticationMode, ""+
		"ssh remote server authentication mode either basic,key,file")

	fs.StringVar(&m.Username, "ssh-username", m.Username, ""+
		"Username for access to remote server when AuthenticationMode set to basic.")

	fs.StringVar(&m.Password, "db-password", m.Password, ""+
		"Password for access to remote server,when AuthenticationMode set to basic.")

	fs.IntVar(&m.SshPort, "ssh-port", m.SshPort, ""+
		"remote server ssh port default 22.")

	fs.StringVar(&m.PrivateKeyFile, "ssh-private-key-file", m.PrivateKeyFile, ""+
		"ssh key private key for access to remote server when AuthenticationMode is set to file.")

	fs.StringVar(&m.PrivateKey, "ssh-private-key", m.PrivateKey, ""+
		"ssh key private key for access to remote server when AuthenticationMode is set to key.")
}

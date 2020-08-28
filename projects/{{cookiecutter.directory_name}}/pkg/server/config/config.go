package config

import (
	"{{cookiecutter.project_slug}}/pkg/client/orm"
	"{{cookiecutter.project_slug}}/pkg/client/redis"
	"{{cookiecutter.project_slug}}/pkg/client/ssh"
	"{{cookiecutter.project_slug}}/pkg/logger"
	"fmt"
	"github.com/spf13/viper"
)

const (
	DefaultConfigurationName = "config"
	DefaultConfigurationPath = "./"
)

var (
	// sharedConfig holds configuration
	sharedConfig *Config

	// shadowConfig contains options from commandline options
	shadowConfig = &Config{}
)

type Config struct {
	OrmOptions   *orm.OrmOptions     `json:"orm,omitempty" yaml:"orm,omitempty" mapstructure:"orm"`
	RedisOptions *redis.RedisOptions `json:"redis,omitempty" yaml:"redis,omitempty" mapstructure:"redis"`
	SshOptions   *ssh.SshOptions     `json:"ssh,omitempty" yaml:"ssh,omitempty" mapstructure:"ssh"`
}

func newConfig() *Config {
	return &Config{
		OrmOptions:   orm.NewOrmOptions(),
		RedisOptions: redis.NewRedisOptions(),
		SshOptions:   ssh.NewOrmOptions(),
	}
}

func Get() *Config {
	return sharedConfig
}

func (c *Config) Apply(conf *Config) {
	shadowConfig = conf

	if conf.RedisOptions != nil {
		conf.RedisOptions.ApplyTo(c.RedisOptions)
	}

	if conf.OrmOptions != nil {
		conf.OrmOptions.ApplyTo(c.OrmOptions)
	}

	if conf.SshOptions != nil {
		conf.SshOptions.ApplyTo(c.SshOptions)
	}
}

func (c *Config) stripEmptyOptions() {
	if c.OrmOptions != nil && c.OrmOptions.Host == "" {
		c.OrmOptions = nil
	}

	if c.RedisOptions != nil && c.RedisOptions.RedisURL == "" {
		c.RedisOptions = nil
	}

	if c.SshOptions != nil && c.SshOptions.Host == "" {
		c.SshOptions = nil
	}
}

// Load loads configuration after setup
func Load() error {
	sharedConfig = newConfig()

	viper.SetConfigName(DefaultConfigurationName)
	viper.AddConfigPath(DefaultConfigurationPath)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Warn(nil, "configuration file not found")
			return nil
		} else {
			panic(fmt.Errorf("error parsing configuration file %s", err))
		}
	}

	conf := newConfig()
	if err := viper.Unmarshal(conf); err != nil {
		logger.Error(nil, "error unmarshal configuration %v", err)
		return err
	} else {
		conf.Apply(shadowConfig)
		sharedConfig = conf
	}

	return nil
}

package orm

import (
	"fmt"
	"time"

	"{{cookiecutter.project_slug}}/pkg/utils/reflectutils"

	"github.com/spf13/pflag"
)

type OrmOptions struct {
	Host                  string        `json:"host,omitempty" yaml:"host" description:"database service host address, for sqlite is filepath"`
	Username              string        `json:"username,omitempty" yaml:"username"`
	Password              string        `json:"-" yaml:"password"`
	DBType                string        `json:"dbType" yaml:"dbType"`
	DBName                string        `json:"dbName" yaml:"dbName"`
	DBPort                string        `json:"dbPort" yaml:"dbPort"`
	MaxIdleConnections    int           `json:"maxIdleConnections,omitempty" yaml:"maxIdleConnections"`
	MaxOpenConnections    int           `json:"maxOpenConnections,omitempty" yaml:"maxOpenConnections"`
	MaxConnectionLifeTime time.Duration `json:"maxConnectionLifeTime,omitempty" yaml:"maxConnectionLifeTime"`
}

func NewOrmOptions() *OrmOptions {
	return &OrmOptions{
		Host:                  "localhost",
		Username:              "",
		Password:              "",
		DBType:                "sqlite3",
		DBName:                "/root/db/{{cookiecutter.project_slug}}.sqlite",
		DBPort:                "",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
	}
}

func (m *OrmOptions) Validate() []error {
	var errors []error

	return errors
}

func (m *OrmOptions) ApplyTo(options *OrmOptions) {
	reflectutils.Override(options, m)
}

func (m *OrmOptions) AddFlags(fs *pflag.FlagSet) {

	fs.StringVar(&m.Host, "db-host", m.Host, ""+
		"database service host address. If left blank, the following related database options will be ignored.")

	fs.StringVar(&m.Username, "db-username", m.Username, ""+
		"Username for access to database service.")

	fs.StringVar(&m.Password, "db-password", m.Password, ""+
		"Password for access to database, should be used pair with password.")

	fs.StringVar(&m.DBType, "db-type", m.DBType, ""+
		"database type, e.g. mysql, sqlite, postgres.")

	fs.StringVar(&m.DBName, "db-name", m.DBName, ""+
		"database name to use")

	fs.StringVar(&m.DBPort, "db-port", m.DBPort, ""+
		"database port to use")

	fs.IntVar(&m.MaxIdleConnections, "db-max-idle-connections", m.MaxOpenConnections, ""+
		"Maximum idle connections allowed to connect to mysql.")

	fs.IntVar(&m.MaxOpenConnections, "db-max-open-connections", m.MaxOpenConnections, ""+
		"Maximum open connections allowed to connect to mysql.")

	fs.DurationVar(&m.MaxConnectionLifeTime, "db-max-connection-life-time", m.MaxConnectionLifeTime, ""+
		"Maximum connection life time allowed to connecto to mysql.")
}

func (m *OrmOptions) GetDBUrl() string {
	switch m.DBType {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true", m.Username, m.Password, m.Host, m.DBName)
	case "sqlite3":
		return m.DBName
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", m.Host, m.DBPort, m.Username, m.DBName, m.Password)
	default:
		return ""
	}
}

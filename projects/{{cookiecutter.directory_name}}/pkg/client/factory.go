package client

import (
	"{{cookiecutter.project_slug}}/pkg/client/orm"
	"{{cookiecutter.project_slug}}/pkg/client/redis"
	"fmt"
	goredis "github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"sync"
)

type ClientSetNotEnabledError struct {
	err error
}

func (e ClientSetNotEnabledError) Error() string {
	return fmt.Sprintf("client set not enabled: %v", e.err)
}

var mutex sync.Mutex

type ClientSetOptions struct {
	ormOptions   *orm.OrmOptions
	redisOptions *redis.RedisOptions
}

func NewClientSetOptions() *ClientSetOptions {
	return &ClientSetOptions{
		ormOptions:   orm.NewOrmOptions(),
		redisOptions: redis.NewRedisOptions(),
	}
}

func (c *ClientSetOptions) SetOrmOptions(options *orm.OrmOptions) *ClientSetOptions {
	c.ormOptions = options
	return c
}

func (c *ClientSetOptions) SetRedisOptions(options *redis.RedisOptions) *ClientSetOptions {
	c.redisOptions = options
	return c
}

// ClientSet provide best of effort service to initialize clients,
// but there is no guarantee to return a valid client instance,
// so do validity check before use
type ClientSet struct {
	csoptions *ClientSetOptions
	stopCh    <-chan struct{}

	database    *orm.DBClient
	redisClient *redis.RedisClient
}

// global clientsets instance
var sharedClientSet *ClientSet

func ClientSets() *ClientSet {
	return sharedClientSet
}

func NewClientSetFactory(c *ClientSetOptions, stopCh <-chan struct{}) *ClientSet {
	sharedClientSet = &ClientSet{csoptions: c, stopCh: stopCh}

	return sharedClientSet
}

func (cs *ClientSet) Redis() (*goredis.Client, error) {
	var err error

	if cs.csoptions.redisOptions == nil || cs.csoptions.redisOptions.RedisURL == "" {
		return nil, ClientSetNotEnabledError{}
	}

	if cs.redisClient != nil {
		return cs.redisClient.Redis(), nil
	} else {
		mutex.Lock()
		defer mutex.Unlock()
		if cs.redisClient == nil {
			cs.redisClient, err = redis.NewRedisClient(cs.csoptions.redisOptions, cs.stopCh)
			if err != nil {
				return nil, err
			}
		}

		return cs.redisClient.Redis(), nil
	}
}

func (cs *ClientSet) Database(modelRegister func(db *gorm.DB)) (*gorm.DB, error) {
	var err error

	if cs.csoptions.ormOptions == nil || cs.csoptions.ormOptions.Host == "" {
		return nil, ClientSetNotEnabledError{}
	}

	if cs.database != nil {
		return cs.database.DB(), nil
	} else {
		mutex.Lock()
		defer mutex.Unlock()
		if cs.database == nil {
			cs.database, err = orm.NewDBClient(cs.csoptions.ormOptions, cs.stopCh, modelRegister)
			if err != nil {
				return nil, err
			}
		}
		return cs.database.DB(), nil
	}
}

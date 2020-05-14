package orm

import (
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/logger"

	"github.com/jinzhu/gorm"

	// gorm存储注入
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DBClient struct {
	db *gorm.DB
}

func NewDBClientOrDie(options *OrmOptions, stopCh <-chan struct{}, modelRegister func(db *gorm.DB)) (*DBClient, error) {
	client, err := NewDBClient(options, stopCh, modelRegister)
	if err != nil {
		panic(err)
	}
	return client, nil
}

func NewDBClient(options *OrmOptions, stopCh <-chan struct{}, modelRegister func(db *gorm.DB)) (*DBClient, error) {
	var dbClient DBClient
	var err error

	dbClient.db, err = gorm.Open(options.DBType, options.GetDBUrl())
	if err != nil {
		logger.Error(nil, "database open error", err)
		return nil, err
	}

	err = dbClient.db.DB().Ping()
	if err != nil {
		_ = dbClient.db.Close()
		logger.Error(nil, "database ping error", err)
		return nil, err
	}

	dbClient.db.DB().SetMaxOpenConns(options.MaxOpenConnections)
	dbClient.db.DB().SetMaxIdleConns(options.MaxIdleConnections)
	dbClient.db.DB().SetConnMaxLifetime(options.MaxConnectionLifeTime)
	dbClient.db.SingularTable(true)
	if modelRegister != nil {
		modelRegister(dbClient.db)
	}
	if stopCh != nil {
		go func() {
			<-stopCh
			if err := dbClient.db.Close(); err != nil {
				logger.Error(nil, err.Error())
			}
		}()
	}

	return &dbClient, nil
}

func (d *DBClient) DB() *gorm.DB {
	return d.db
}

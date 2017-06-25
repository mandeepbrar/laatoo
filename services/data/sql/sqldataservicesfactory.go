package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "google.golang.org/appengine/cloudsql"
	// import _ "github.com/jinzhu/gorm/dialects/sqlite"
	// import _ "github.com/jinzhu/gorm/dialects/mssql

	"laatoo/sdk/config"
	"laatoo/sdk/core"
	"laatoo/sdk/errors"
	//"laatoosdk/log"
)

type sqlDataServicesFactory struct {
	vendor           string
	connectionString string
	cache            bool
}

const (
	CONF_SQL_CONNECTIONSTRING = "connectionstring"
	CONF_SQL_VENDOR           = "vendor"
	CONF_SQL_SERVICES         = "sql_services"
)

func Manifest() []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: CONF_SQL_SERVICES, Object: sqlDataServicesFactory{}}}
}

func (sf *sqlDataServicesFactory) Initialize(ctx core.ServerContext, conf config.Config) error {
	connectionString, ok := conf.GetString(CONF_SQL_CONNECTIONSTRING)
	if !ok {
		return errors.MissingConf(ctx, CONF_SQL_CONNECTIONSTRING)
	}
	vendor, ok := conf.GetString(CONF_SQL_VENDOR)
	if !ok {
		return errors.MissingConf(ctx, CONF_SQL_VENDOR)
	}
	sf.connectionString = connectionString
	sf.vendor = vendor
	return nil
}

//Create the services configured for factory.
func (sf *sqlDataServicesFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	return newSqlDataService(ctx, name, sf)
}

//The services start serving when this method is called
func (sf *sqlDataServicesFactory) Start(ctx core.ServerContext) error {
	return nil
}

package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "google.golang.org/appengine/cloudsql"
	// import _ "github.com/jinzhu/gorm/dialects/sqlite"
	// import _ "github.com/jinzhu/gorm/dialects/mssql

	"laatoo/sdk/config"
	"laatoo/sdk/core"
	//"laatoosdk/log"
)

type sqlDataServicesFactory struct {
	core.ServiceFactory
	vendor           string
	connectionString string
	cache            bool
}

const (
	CONF_SQL_CONNECTIONSTRING = "sqlconnectionstring"
	CONF_SQL_VENDOR           = "sqlvendor"
	SQL_FACTORY               = "sql_services"
)

func Manifest(provider core.MetaDataProvider) []core.PluginComponent {
	return []core.PluginComponent{core.PluginComponent{Name: SQL_FACTORY, Object: sqlDataServicesFactory{}}}
}

/*func (sf *sqlDataServicesFactory) Initialize(ctx core.ServerContext) error {
	sf.AddStringConfiguration(ctx, CONF_SQL_CONNECTIONSTRING)
	sf.AddStringConfiguration(ctx, CONF_SQL_VENDOR)
	return nil
}*/

//Create the services configured for factory.
func (sf *sqlDataServicesFactory) CreateService(ctx core.ServerContext, name string, method string, conf config.Config) (core.Service, error) {
	return newSqlDataService(ctx, name, sf)
}

//The services start serving when this method is called
func (sf *sqlDataServicesFactory) Start(ctx core.ServerContext) error {
	connectionString, _ := sf.GetStringConfiguration(ctx, CONF_SQL_CONNECTIONSTRING)
	vendor, _ := sf.GetStringConfiguration(ctx, CONF_SQL_VENDOR)

	sf.connectionString = connectionString
	sf.vendor = vendor
	return nil
}

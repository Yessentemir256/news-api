// config/config.go
package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}

func GetServerAddress() string {
	return viper.GetString("server.address")
}

func GetDatabaseDSN() string {
	// Define your database DSN here, for example:
	// return "postgres://user:password@localhost/dbname"
	return "your_database_dsn_here"
}

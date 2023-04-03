package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type DBConfig struct {
// 	DBDriver string `mapstructure:"DB_DRIVER"`
// 	dbUser   string `mapstructure:"DB_USER"`
// 	dbPort   string `mapstructure:"DB_PORT"`
// 	dbPass   string `mapstructure:"DB_PASSWORD"`
// 	dbHost   string `mapstructure:"DB_HOST"`
// 	dbName   string `mapstructure:"DB_NAME"`
// }

func ConnectDB() (*gorm.DB, error) {

	// Load configuration from a config file or environment variables using Viper
	viper.SetConfigFile(".env") // Or set the file path to your config file
	viper.SetEnvPrefix("mysql")
	viper.AutomaticEnv()

	// Set default values for configuration parameters
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "3306")
	viper.SetDefault("user", "root")
	viper.SetDefault("password", "12345678")
	viper.SetDefault("db_name", "task_manager")

	host := viper.GetString("host")
	port := viper.GetString("port")
	user := viper.GetString("user")
	password := viper.GetString("password")
	dbName := viper.GetString("db_name")

	// Create the connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, dbName)

	// Connect to the database using GORM
	db, err := gorm.Open(mysql.Open(connectionString))
	if err != nil {
		return nil, err
	}

	// Return the database connection
	return db, nil
}

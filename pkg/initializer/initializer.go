package initializer

import (
	"fmt"
	"log"

	"github.com/floriwan/govam/models"
	"github.com/floriwan/govam/pkg/database"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBHost     string `mapstructure:"DATABASE_HOST"`
	DBPort     string `mapstructure:"DATABASE_PORT"`
	DBName     string `mapstructure:"DATABASE_DB"`
	DBUser     string `mapstructure:"DATABASE_USER"`
	DBPassword string `mapstructure:"DATABASE_PASSWORD"`
	DBType     string `mapstructure:"DATABASE_TYPE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func ConnectDB(c Config) error {

	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
	database.Instance, err = gorm.Open(mysql.Open(connectionString),
		&gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func Migrate() {
	database.Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}

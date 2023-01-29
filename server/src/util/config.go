package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment       string `mapstructure:"ENVIRONMENT"`
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	MigrationURL      string `mapstructure:"MIGRATION_URL"`
	RedisAddress      string `mapstructure:"REDIS_ADDRESS"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
	//MONGODB_URI
	MongoDBURI string `mapstructure:"MONGODB_URI"`
	// FIREBASE_DATABASE_URL
	// FIREBASE_SERVICE_ACCOUNT_KEY_JSON_PATH
	FirebaseDatabaseURL               string `mapstructure:"FIREBASE_DATABASE_URL"`
	FirebaseServiceAccountKeyJSONPath string `mapstructure:"FIREBASE_SERVICE_ACCOUNT_KEY_JSON_PATH"`
	// NEO4J_PASSWORD
	// NEO4J_USER
	// NEO4J_HOST
	Neo4jPassword string `mapstructure:"NEO4J_PASSWORD"`
	Neo4jUser     string `mapstructure:"NEO4J_USER"`
	Neo4jHost     string `mapstructure:"NEO4J_HOST"`
	// POSTGRES_HOST
	// POSTGRES_PORT
	// POSTGRES_USER
	// POSTGRES_PASSWORD
	// POSTGRES_DATABASE
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDatabase string `mapstructure:"POSTGRES_DATABASE"`
	// REDIS_DB
	// REDIS_HOST
	// REDIS_PORT
	// REDIS_PASSWORD
	// REDIS_USER
	RedisDB       string `mapstructure:"REDIS_DB"`
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisUser     string `mapstructure:"REDIS_USER"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

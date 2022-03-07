package mongo

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnConfig encapsulates the fields and options of a mongoDB connection.
type ConnConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	SSLPath  string
	Local    bool
	NoTLS    bool
}

// NewClient creates and returns a new mongoDB client for the given
// connection string.
func NewClient(connString string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connString))
	if err != nil {
		panic(err)
	}

	return client
}

// NewClientFromEnv this factory function bypasses params and pulls them
// right from the environment. Useful when running in many environments.
func NewClientFromEnv() *mongo.Client {
	var connString string

	cfg := ParseEnv()
	if cfg.Local {
		connString = GetLocalConnString(cfg)
	} else if cfg.NoTLS {
		connString = GetConnString(cfg)
	} else {
		connString = GetTLSConnString(cfg)
	}

	return NewClient(connString)
}

// ParseEnv parses the mongo related environment variables and stores them in
// a ConnConfig struct.
func ParseEnv() ConnConfig {
	return ConnConfig{
		User:     os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		SSLPath:  os.Getenv("MONGO_SSL_PATH"),
		Local:    os.Getenv("MONGO_LOCAL") == "true",
		NoTLS:    os.Getenv("MONGO_NO_TLS") == "true",
	}
}

// GetTLSConnString use this to connect to a mongodb cluster using basic auth
// and a TLS certificate.
func GetTLSConnString(cfg ConnConfig) string {
	return fmt.Sprintf(
		"mongodb+srv://%v:%v@%v/?tls=true&tlsCAFile=%v",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.SSLPath,
	)
}

// GetConnString use this to connect to a mongodb cluster using
// basic auth only (no TLS)
func GetConnString(cfg ConnConfig) string {
	return fmt.Sprintf(
		"mongodb+srv://%v:%v@%v",
		cfg.User,
		cfg.Password,
		cfg.Host,
	)
}

// GetLocalConnString use this for connecting to local (non-cluster) instances
func GetLocalConnString(cfg ConnConfig) string {
	return fmt.Sprintf(
		"mongodb://%v:%v@%v:%v",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)
}

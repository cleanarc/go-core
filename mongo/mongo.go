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
}

// NewClient accepts a context and a ConnConfig. Returns a new
// mongoDB client configured with the passed in params.
// if ctx is nil, context.TODO will be used.
func NewClient(ctx context.Context, cfg ConnConfig) *mongo.Client {
	// Set up the context
	if ctx == nil {
		ctx = context.TODO()
	}

	// Parse the config
	connString := newConnString(cfg)

	// Build the client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		panic(err)
	}

	return client
}

// ParseEnv parses the environment for mongoDB configurations. Returns a
// ConnConfig which encapsulates the results. Set the parsed variables in your
// environment when running your application, and this convenience function will
// pick them up. Parsed Variables: MONGO_USER, MONGO_PASSWORD, MONGO_HOST
func ParseEnv() ConnConfig {
	return ConnConfig{
		User:     os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
	}
}

func newConnString(cfg ConnConfig) string {
	return fmt.Sprintf(
		"mongodb://%v:%v@%v:%v",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)
}

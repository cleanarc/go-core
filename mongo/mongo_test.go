package mongo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEnv(t *testing.T) {
	// expected output
	expected := ConnConfig{
		User:     "root",
		Password: "password",
		Host:     "localhost",
		Port:     "27017",
		SSLPath:  "/etc/mongo/ssl",
		Local:    true,
		NoTLS:    true,
	}

	// Setup env
	_ = os.Setenv("MONGO_USER", "root")
	_ = os.Setenv("MONGO_PASSWORD", "password")
	_ = os.Setenv("MONGO_HOST", "localhost")
	_ = os.Setenv("MONGO_PORT", "27017")
	_ = os.Setenv("MONGO_SSL_PATH", "/etc/mongo/ssl")
	_ = os.Setenv("MONGO_LOCAL", "true")
	_ = os.Setenv("MONGO_NO_TLS", "true")

	assert.Equal(t, expected, ParseEnv())
}

func TestGetTLSConnString(t *testing.T) {
	tests := []struct {
		name     string
		cfg      ConnConfig
		expected string
	}{
		{
			name: "remote-tls-uri",
			cfg: ConnConfig{
				User:     "user",
				Password: "password",
				Host:     "wcbr-mongodb-1e9322e2.mongo.ondigitalocean.com",
				SSLPath:  "/etc/mongo/ssl",
			},
			expected: "mongodb+srv://user:password@wcbr-mongodb-1e9322e2.mongo.ondigitalocean.com/?tls=true&tlsCAFile=/etc/mongo/ssl",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, GetTLSConnString(test.cfg))
		})
	}
}

func TestGetConnString(t *testing.T) {
	tests := []struct {
		name     string
		cfg      ConnConfig
		expected string
	}{
		{
			name: "remote-no-tls-uri",
			cfg: ConnConfig{
				User:     "user",
				Password: "password",
				Host:     "wcbr-mongodb-1e9322e2.mongo.ondigitalocean.com",
				NoTLS:    true,
			},
			expected: "mongodb+srv://user:password@wcbr-mongodb-1e9322e2.mongo.ondigitalocean.com",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, GetConnString(test.cfg))
		})
	}
}

func TestGetLocalConnString(t *testing.T) {
	tests := []struct {
		name     string
		cfg      ConnConfig
		expected string
	}{
		{
			name: "localhost",
			cfg: ConnConfig{
				User:     "user",
				Password: "password",
				Host:     "localhost",
				Port:     "27017",
				Local:    true,
			},
			expected: "mongodb://user:password@localhost:27017",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, GetLocalConnString(test.cfg))
		})
	}
}

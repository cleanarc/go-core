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
	}

	// Setup env
	_ = os.Setenv("MONGO_USER", "root")
	_ = os.Setenv("MONGO_PASSWORD", "password")
	_ = os.Setenv("MONGO_HOST", "localhost")
	_ = os.Setenv("MONGO_PORT", "27017")

	assert.Equal(t, expected, ParseEnv())
}

func Test_newConnString(t *testing.T) {
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
			},
			expected: "mongodb://user:password@localhost:27017",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, newConnString(test.cfg))
		})
	}
}

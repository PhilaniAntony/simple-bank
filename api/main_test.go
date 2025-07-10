package api

import (
	"os"
	"testing"
	"time"

	db "github.com/PhilaniAntony/simplebank/db/sqlc"
	"github.com/PhilaniAntony/simplebank/util"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

func newTestServer(t *testing.T, store *db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   getEnvOrDefault("TOKEN_SYMMETRIC_KEY", util.RandomString(32)),
		AccessTokenDuration: getDurationOrDefault("ACCESS_TOKEN_DURATION", time.Minute),
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func getEnvOrDefault(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

func getDurationOrDefault(key string, defaultVal time.Duration) time.Duration {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	duration, err := time.ParseDuration(val)
	if err != nil {
		return defaultVal
	}
	return duration
}

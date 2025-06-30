package api

import (
	"testing"
	"time"

	db "github.com/PhilaniAntony/simplebank/db/sqlc"
	"github.com/PhilaniAntony/simplebank/util"
)

func newTestServer(t *testing.T, store *db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	if err != nil {
		return nil
	}

	return server
}

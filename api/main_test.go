package api

import (
	"testing"

	db "github.com/PhilaniAntony/simplebank/db/sqlc"
	"github.com/PhilaniAntony/simplebank/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store *db.Store) *Server {
	config, err := util.LoadConfig(".")
	require.NoError(t, err)

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

package api

import (
	"os"
	"testing"
	"time"

	db "github.com/copter01252/simplebank/db/sqlc"
	"github.com/copter01252/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServe(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(t *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(t.Run())
}

package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	db "github.com/techscholl/simplebank/db/sqlc"
	"github.com/techscholl/simplebank/util"
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

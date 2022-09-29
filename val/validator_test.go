package val

import (
	"testing"

	"github.com/copter01252/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestValidateString(t *testing.T) {
	string := util.RandomString(10)
	err := ValidateString(string, 3, 10)
	require.NoError(t, err)
}

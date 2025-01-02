package bitaxe

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	baseAddress = strings.TrimSpace(os.Getenv("BITAXE_BASE_ADDRESS"))
)

func TestClient_SystemInfo(t *testing.T) {
	c := testClient(t)

	ctx := context.Background()
	info, err := c.SystemInfo(ctx)
	require.NoError(t, err)

	require.NotEmpty(t, info.BestDiff)
}

func testClient(t *testing.T) Client {
	t.Helper()

	if baseAddress == "" {
		t.Skip("BITAXE_BASE_ADDRESS is empty")
	}

	return NewClient(baseAddress, nil)
}

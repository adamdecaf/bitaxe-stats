package bitaxe

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	defaultBaseAddress = strings.TrimSpace(os.Getenv("BITAXE_BASE_ADDRESS"))
)

func TestClient_SystemInfo(t *testing.T) {
	c := testClient(t)

	ctx := context.Background()
	info, err := c.SystemInfo(ctx, baseAddress(t))
	require.NoError(t, err)

	require.NotEmpty(t, info.BestDiff)
}

func testClient(t *testing.T) Client {
	t.Helper()

	if testing.Short() {
		t.Skip("-short flag provided")
	}

	return NewClient(nil)
}

func baseAddress(t *testing.T) string {
	t.Helper()

	if defaultBaseAddress == "" {
		t.Skip("BITAXE_BASE_ADDRESS is empty")
	}

	return defaultBaseAddress
}

package bitaxe

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSystemInfo(t *testing.T) {
	fd, err := os.Open(filepath.Join("..", "..", "testdata", "api", "system-info.json"))
	require.NoError(t, err)

	t.Cleanup(func() { fd.Close() })

	dec := json.NewDecoder(fd)
	dec.DisallowUnknownFields()

	var info SystemInfo
	err = dec.Decode(&info)
	require.NoError(t, err)

	require.InDelta(t, 16.78509521484375, info.Power, 0.01)
	require.InDelta(t, 5179.6875, info.Voltage, 0.01)
	require.InDelta(t, 10296.875, info.Current, 0.01)
	require.InDelta(t, 52.875, info.Temp, 0.01)
	require.Equal(t, 52, info.VrTemp)
	require.InDelta(t, 1040.3401762854942, info.HashRate, 0.01)
	require.Equal(t, "115M", info.BestDiff)
	require.Equal(t, "115M", info.BestSessionDiff)
	require.Equal(t, 0, info.IsUsingFallbackStratum)
	require.Equal(t, 146916, info.FreeHeap)
	require.Equal(t, 1150, info.CoreVoltage)
	require.Equal(t, 1133, info.CoreVoltageActual)
	require.Equal(t, 525, info.Frequency)
	require.Equal(t, "Home Wifi", info.SSID)
	require.Equal(t, "E4:B0:63:92:01:04", info.MacAddr)
	require.Equal(t, "bitaxe_0123", info.Hostname)
	require.Equal(t, "Connected!", info.WifiStatus)
	require.Equal(t, 26969, info.SharesAccepted)
	require.Equal(t, 207, info.SharesRejected)
	require.Equal(t, 257902, info.UptimeSeconds)
	require.Equal(t, 1, info.AsicCount)
	require.Equal(t, 2040, info.SmallCoreCount)
	require.Equal(t, "BM1370", info.ASICModel)
	require.Equal(t, "public-pool.io", info.StratumURL)
	require.Equal(t, "solo.ckpool.org", info.FallbackStratumURL)
	require.Equal(t, 21496, info.StratumPort)
	require.Equal(t, 3333, info.FallbackStratumPort)
	require.Equal(t, "address.bitaxe_0123", info.StratumUser)
	require.Equal(t, "address.bitaxe_0123", info.FallbackStratumUser)
	require.Equal(t, "v2.4.0", info.Version)
	require.Equal(t, "601", info.BoardVersion)
	require.Equal(t, "factory", info.RunningPartition)
	require.Equal(t, 1, info.Flipscreen)
	require.Equal(t, 0, info.OverheatMode)
	require.Equal(t, 0, info.Invertscreen)
	require.Equal(t, 1, info.Invertfanpolarity)
	require.Equal(t, 1, info.Autofanspeed)
	require.Equal(t, 52, info.Fanspeed)
	require.Equal(t, 5362, info.Fanrpm)
}

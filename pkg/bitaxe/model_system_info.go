package bitaxe

type SystemInfo struct {
	Power                  float64 `json:"power"`
	Voltage                float64 `json:"voltage"`
	Current                float64 `json:"current"`
	Temp                   float64 `json:"temp"`
	VrTemp                 int     `json:"vrTemp"`
	HashRate               float64 `json:"hashRate"`
	BestDiff               string  `json:"bestDiff"`
	BestSessionDiff        string  `json:"bestSessionDiff"`
	IsUsingFallbackStratum int     `json:"isUsingFallbackStratum"`
	FreeHeap               int     `json:"freeHeap"`
	CoreVoltage            int     `json:"coreVoltage"`
	CoreVoltageActual      int     `json:"coreVoltageActual"`
	Frequency              int     `json:"frequency"`
	SSID                   string  `json:"ssid"`
	MacAddr                string  `json:"macAddr"`
	Hostname               string  `json:"hostname"`
	WifiStatus             string  `json:"wifiStatus"`
	SharesAccepted         int     `json:"sharesAccepted"`
	SharesRejected         int     `json:"sharesRejected"`
	UptimeSeconds          int     `json:"uptimeSeconds"`
	AsicCount              int     `json:"asicCount"`
	SmallCoreCount         int     `json:"smallCoreCount"`
	ASICModel              string  `json:"ASICModel"`
	StratumURL             string  `json:"stratumURL"`
	FallbackStratumURL     string  `json:"fallbackStratumURL"`
	StratumPort            int     `json:"stratumPort"`
	FallbackStratumPort    int     `json:"fallbackStratumPort"`
	StratumUser            string  `json:"stratumUser"`
	FallbackStratumUser    string  `json:"fallbackStratumUser"`
	Version                string  `json:"version"`
	BoardVersion           string  `json:"boardVersion"`
	RunningPartition       string  `json:"runningPartition"`
	Flipscreen             int     `json:"flipscreen"`
	OverheatMode           int     `json:"overheat_mode"`
	Invertscreen           int     `json:"invertscreen"`
	Invertfanpolarity      int     `json:"invertfanpolarity"`
	Autofanspeed           int     `json:"autofanspeed"`
	Fanspeed               int     `json:"fanspeed"`
	Fanrpm                 int     `json:"fanrpm"`
}

package bitaxe

type SystemInfo struct {
	Power                  float64                `json:"power"`
	Voltage                float64                `json:"voltage"`
	Current                float64                `json:"current"`
	Temp                   float64                `json:"temp"`
	VrTemp                 int                    `json:"vrTemp"`
	MaxPower               int                    `json:"maxPower"`
	NominalVoltage         int                    `json:"nominalVoltage"`
	HashRate               float64                `json:"hashRate"`
	ExpectedHashrate       int                    `json:"expectedHashrate"`
	BestDiff               string                 `json:"bestDiff"`
	BestSessionDiff        string                 `json:"bestSessionDiff"`
	StratumDiff            int                    `json:"stratumDiff"`
	IsUsingFallbackStratum int                    `json:"isUsingFallbackStratum"`
	IsPSRAMAvailable       int                    `json:"isPSRAMAvailable"`
	FreeHeap               int                    `json:"freeHeap"`
	CoreVoltage            int                    `json:"coreVoltage"`
	CoreVoltageActual      int                    `json:"coreVoltageActual"`
	Frequency              int                    `json:"frequency"`
	Ssid                   string                 `json:"ssid"`
	MacAddr                string                 `json:"macAddr"`
	Hostname               string                 `json:"hostname"`
	WifiStatus             string                 `json:"wifiStatus"`
	WifiRSSI               int                    `json:"wifiRSSI"`
	ApEnabled              int                    `json:"apEnabled"`
	SharesAccepted         int                    `json:"sharesAccepted"`
	SharesRejected         int                    `json:"sharesRejected"`
	SharesRejectedReasons  []SharesRejectedReason `json:"sharesRejectedReasons"`
	UptimeSeconds          int                    `json:"uptimeSeconds"`
	AsicCount              int                    `json:"asicCount"`
	SmallCoreCount         int                    `json:"smallCoreCount"`
	ASICModel              string                 `json:"ASICModel"`
	StratumURL             string                 `json:"stratumURL"`
	FallbackStratumURL     string                 `json:"fallbackStratumURL"`
	StratumPort            int                    `json:"stratumPort"`
	FallbackStratumPort    int                    `json:"fallbackStratumPort"`
	StratumUser            string                 `json:"stratumUser"`
	FallbackStratumUser    string                 `json:"fallbackStratumUser"`
	Version                string                 `json:"version"`
	IdfVersion             string                 `json:"idfVersion"`
	BoardVersion           string                 `json:"boardVersion"`
	RunningPartition       string                 `json:"runningPartition"`
	OverheatMode           int                    `json:"overheat_mode"`
	OverclockEnabled       int                    `json:"overclockEnabled"`
	Display                string                 `json:"display"`
	FlipScreen             int                    `json:"flipscreen"`
	InvertScreen           int                    `json:"invertscreen"`
	DisplayTimeout         int                    `json:"displayTimeout"`
	AutoFanSpeed           int                    `json:"autofanspeed"`
	FanSpeed               int                    `json:"fanspeed"`
	TempTarget             int                    `json:"temptarget"`
	FanRPM                 int                    `json:"fanrpm"`
	StatsLimit             int                    `json:"statsLimit"`
	StatsDuration          int                    `json:"statsDuration"`
}

type SharesRejectedReason struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

package nhclient

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"math"
)

type device struct {
	ID         string
	Name       string
	DeviceType *struct {
		EnumName    deviceType
		Description string
	}
	Status *struct {
		EnumName    deviceStatus
		Description string
	}
	Temperature                    float64
	Load                           float64
	RevolutionsPerMinute           float64
	RevolutionsPerMinutePercentage float64
	PowerMode                      *struct {
		EnumName    powerMode
		Description string
	}
	PowerUsage float64
	Speeds     []*struct {
		Algorithm     algorithm
		Title         string
		Speed         string //Wrong type on doc
		DisplaySuffix string
	}
	Intensity *struct {
		EnumName    intensity
		Description string
	}
	NHQM string
}

func (d *device) GetHotspotTemperature() float64 {
	return d.Temperature / 65536
}

func (d *device) GetCoreTemperature() float64 {
	return math.Mod(d.Temperature, 65536)
}

type stats struct {
	StatsTime int64
	Market    string //TODO: Enum
	Algorithm *struct {
		EnumName    string //TODO: Enum
		Description string
	}
	UnpaidAmount             decimal.Decimal
	Difficulty               float64
	ProxyID                  int64
	TimeConnected            int64
	XNSub                    bool
	SpeedAccepted            float64
	SpeedRejectedR1Target    float64
	SpeedRejectedR2Stale     float64
	SpeedRejectedR3Duplicate float64
	SpeedRejectedR4NTime     float64
	SpeedRejectedR5Other     float64
	SpeedRejectedTotal       float64
	Profitability            float64
}

type rig struct {
	Id                 string `json:"rigId"`
	RigType            string `json:"type"` //TODO: Enum
	Name               string
	StatusTime         int64
	JoinTime           int64
	MinerStatus        string //TODO: Enum
	GroupName          string
	UnpaidAmount       decimal.Decimal
	Notifications      []string //TODO: Enum
	SoftwareVersions   string
	Devices            []*device
	CPUMiningEnabled   bool
	CPUExists          bool
	Stats              []*stats
	Profitability      decimal.Decimal
	LocalProfitability decimal.Decimal
	RigPowerMode       string //TODO: Enum
}

type deviceType uint

const (
	DeviceTypeUnknown deviceType = iota
	DeviceTypeNvidia
	DeviceTypeAMD
	DeviceTypeCPU
)

func (e deviceType) String() string {
	return map[deviceType]string{
		DeviceTypeUnknown: "deviceStatusUnknown",
		DeviceTypeNvidia:  "NVIDIA",
		DeviceTypeAMD:     "AMD",
		DeviceTypeCPU:     "CPU",
	}[e]
}

func (e *deviceType) FromString(es string) deviceType {
	return map[string]deviceType{
		"deviceStatusUnknown": DeviceTypeUnknown,
		"NVIDIA":              DeviceTypeNvidia,
		"AMD":                 DeviceTypeAMD,
		"CPU":                 DeviceTypeCPU,
	}[es]
}

func (e deviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *deviceType) UnmarshalJSON(bytes []byte) error {
	var s string
	err := json.Unmarshal(bytes, &s)
	if err != nil {
		return err
	}
	*e = e.FromString(s)
	return nil
}

//--------------------------//

type deviceStatus uint

const (
	DeviceStatusUnknown deviceStatus = iota
	DeviceStatusDisabled
	DeviceStatusInactive
	DeviceStatusMining
	DeviceStatusBenchmarking
	DeviceStatusError
	DeviceStatusPending
	DeviceStatusOffline
)

func (e deviceStatus) String() string {
	return map[deviceStatus]string{
		DeviceStatusUnknown:      "UNKNOWN",
		DeviceStatusDisabled:     "DISABLED",
		DeviceStatusInactive:     "INACTIVE",
		DeviceStatusMining:       "MINING",
		DeviceStatusBenchmarking: "BENCHMARKING",
		DeviceStatusError:        "ERROR",
		DeviceStatusPending:      "PENDING",
		DeviceStatusOffline:      "OFFLINE",
	}[e]
}

func (e *deviceStatus) FromString(es string) deviceStatus {
	return map[string]deviceStatus{
		"UNKNOWN":      DeviceStatusUnknown,
		"DISABLED":     DeviceStatusDisabled,
		"INACTIVE":     DeviceStatusInactive,
		"MINING":       DeviceStatusMining,
		"BENCHMARKING": DeviceStatusBenchmarking,
		"ERROR":        DeviceStatusError,
		"PENDING":      DeviceStatusPending,
		"OFFLINE":      DeviceStatusOffline,
	}[es]
}

func (e deviceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *deviceStatus) UnmarshalJSON(bytes []byte) error {
	var s string
	err := json.Unmarshal(bytes, &s)
	if err != nil {
		return err
	}
	*e = e.FromString(s)
	return nil
}

//-------------------------//

type powerMode uint

const (
	PowerModeUnknown powerMode = iota
	PowerModeLow
	PowerModeMedium
	PowerModeHigh
	PowerModeMixed
)

func (e powerMode) String() string {
	return map[powerMode]string{
		PowerModeUnknown: "UNKNOWN",
		PowerModeLow:     "LOW",
		PowerModeMedium:  "MEDIUM",
		PowerModeHigh:    "HIGH",
		PowerModeMixed:   "MIXED",
	}[e]
}

func (e *powerMode) FromString(es string) powerMode {
	return map[string]powerMode{
		"UNKNOWN": PowerModeUnknown,
		"LOW":     PowerModeLow,
		"MEDIUM":  PowerModeMedium,
		"HIGH":    PowerModeHigh,
		"MIXED":   PowerModeMixed,
	}[es]
}

func (e powerMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *powerMode) UnmarshalJSON(bytes []byte) error {
	var s string
	err := json.Unmarshal(bytes, &s)
	if err != nil {
		return err
	}
	*e = e.FromString(s)
	return nil
}

//------------------------//

type intensity uint

const (
	IntensityUnknown intensity = iota
	IntensityLow
	IntensityHigh
)

func (e intensity) String() string {
	return map[intensity]string{
		IntensityUnknown: "UNKNOWN",
		IntensityLow:     "LOW",
		IntensityHigh:    "HIGH",
	}[e]
}

func (e *intensity) FromString(es string) intensity {
	return map[string]intensity{
		"UNKNOWN": IntensityUnknown,
		"LOW":     IntensityLow,
		"HIGH":    IntensityHigh,
	}[es]
}

func (e intensity) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *intensity) UnmarshalJSON(bytes []byte) error {
	var s string
	err := json.Unmarshal(bytes, &s)
	if err != nil {
		return err
	}
	*e = e.FromString(s)
	return nil
}

//-------------------------------//

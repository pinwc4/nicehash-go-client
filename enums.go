package nhclient

import (
	"encoding/json"
)

type deviceType uint

const (
	DeviceTypeUnknown deviceType = iota
	DeviceTypeNvidia
	DeviceTypeAMD
	DeviceTypeCPU
)

func (e deviceType) String() string {
	return map[deviceType]string{
		DeviceTypeUnknown: "UNKNOWN",
		DeviceTypeNvidia:  "NVIDIA",
		DeviceTypeAMD:     "AMD",
		DeviceTypeCPU:     "CPU",
	}[e]
}

func (e *deviceType) FromString(es string) deviceType {
	return map[string]deviceType{
		"UNKNOWN": DeviceTypeUnknown,
		"NVIDIA":  DeviceTypeNvidia,
		"AMD":     DeviceTypeAMD,
		"CPU":     DeviceTypeCPU,
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

type powerMode uint

const (
	PowerModeUnknown powerMode = iota
	PowerModeDisabled
	PowerModeInactive
	PowerModeMining
	PowerModeBenchmarking
	PowerModeError
	PowerModePending
	PowerModeOffline
)

func (e powerMode) String() string {
	return map[powerMode]string{
		PowerModeUnknown:      "UNKNOWN",
		PowerModeDisabled:     "DISABLED",
		PowerModeInactive:     "INACTIVE",
		PowerModeMining:       "MINING",
		PowerModeBenchmarking: "BENCHMARKING",
		PowerModeError:        "ERROR",
		PowerModePending:      "PENDING",
		PowerModeOffline:      "OFFLINE",
	}[e]
}

func (e *powerMode) FromString(es string) powerMode {
	return map[string]powerMode{
		"UNKNOWN":      PowerModeUnknown,
		"DISABLED":     PowerModeDisabled,
		"INACTIVE":     PowerModeInactive,
		"MINING":       PowerModeMining,
		"BENCHMARKING": PowerModeBenchmarking,
		"ERROR":        PowerModeError,
		"PENDING":      PowerModePending,
		"OFFLINE":      PowerModeOffline,
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

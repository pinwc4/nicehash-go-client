package nhclient

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/shopspring/decimal"
	"math"
)

const BaseURL = "https://api2.nicehash.com"

type device struct {
	ID          string
	Name        string
	Temperature float64
	Load        float64
}

func (d *device) GetMemoryTemperature() float64 {
	return d.Temperature / 65536
}

func (d *device) GetCoreTemperature() float64 {
	return math.Mod(d.Temperature, 65536)
}

type stats struct {
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

type mining struct {
	httpClient *resty.Client
}

//GetAddress get the mining address.
func (m *mining) GetAddress() (string, error) {
	type response struct {
		Address string
	}

	resp, err := m.httpClient.R().
		SetResult(&response{}).
		Get(BaseURL + "/main/api/v2/mining/miningAddress")

	if err != nil {
		return "", err
	}

	r := resp.Result().(*response)
	return r.Address, nil
}

//GetAlgos list mining algos with basic statistics for organization (and for rig id if specified).
func (m *mining) GetAlgos(rigId *string) {
	//TODO: Implement method
}

//GetRigsGroups list groups with list of rigs in the groups.
//When extendedResponse is set to true, response contains number of total and active devices for each rig and group.
func (m *mining) GetRigsGroups() {
	//TODO: Implement method
}

//GetRigStats get statistical streams for selected rig.
func (m *mining) GetRigStats() {
	//TODO: Implement method
}

//GetRigStatisticsByAlgo get statistical streams for selected rig and selected algorithm.
func (m *mining) GetRigStatisticsByAlgo(algo string) {
	//TODO: Implement method
}

//GetRigDetails get mining rig detailed information for selected rig.
func (m *mining) GetRigDetails(rigId string) (*rig, error) {
	resp, err := m.httpClient.R().
		SetResult(&rig{}).
		Get(fmt.Sprintf("%s/main/api/v2/mining/rig2/%s", BaseURL, rigId))

	//TODO: Improve error handling
	if err != nil {
		return nil, err
	}

	//TODO: Improve status check
	//if resp.StatusCode() != 200 {
		return nil, errors.New("unauthorized")
	}

	return resp.Result().(*rig), nil
}

//GetActiveWorkers get a list of active worker.
func (m *mining) GetActiveWorkers() {
	//TODO: Implement method
}

//GetPayouts get list of payouts.
func (m *mining) GetPayouts() {
	//TODO: Implement method
}

// GetAllRigsStatistics get statistical streams for all mining rigs.
func (m *mining) GetAllRigsStatistics() {
	//TODO: Implement method
}

//GetAllRigsStatisticsByAlgo get statistical streams for all mining rigs for selected algorithm. Algorithm code can be found in buy info endpoint.
func (m *mining) GetAllRigsStatisticsByAlgo(algo string) {
	//TODO: Implement method
}

//GetAllRigs list rigs and their statuses.
//Path parameter filters rigs by group.
//When path is empty, rigs from root group are returned.
//Rigs can be sorted according to sort parameter.
func (m *mining) GetAllRigs(algo string) {
	//TODO: Implement method
}

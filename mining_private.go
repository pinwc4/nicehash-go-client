package nhclient

import (
	"github.com/shopspring/decimal"
	"math"
)

type miningPrivate struct {
	client *client
}

type device struct {
	ID         string
	Name       string
	DeviceType *struct {
		EnumName    string //Todo: enum
		Description string
	}
	Temperature                    float64
	Load                           float64
	RevolutionsPerMinute           float64
	RevolutionsPerMinutePercentage float64
	PowerMode                      *struct {
		EnumName    string //Todo: Enum
		Description string
	}
	PowerUsage float64
	Speeds     []*struct {
		Algorithm     string //Todo: Enum
		Title         string
		Speed         string //TODO: Check float
		DisplaySuffix string
	}
	Intensity *struct {
		EnumName    string //Todo: Enum
		Description string
	}
	NHQM string
}

func (d *device) GetMemoryTemperature() float64 {
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

//GetAddress get the miningPrivate address.
func (m *miningPrivate) GetAddress() (string, error) {
	m.client.doRequest()

	return "", nil

	/*type response struct {
		Address string
	}

	resp, err := m.httpClient.R().
		SetResult(&response{}).
		Get(ProdUrl + "/main/api/v2/miningPrivate/miningAddress")

	if err != nil {
		return "", err
	}

	r := resp.Result().(*response)
	return r.Address, nil*/
}

//GetAlgos list miningPrivate algos with basic statistics for organization (and for rig id if specified).
func (m *miningPrivate) GetAlgos(rigId *string) {
	//TODO: Implement method
}

//GetRigsGroups list groups with list of rigs in the groups.
//When extendedResponse is set to true, response contains number of total and active devices for each rig and group.
func (m *miningPrivate) GetRigsGroups() {
	//TODO: Implement method
}

//GetRigStats get statistical streams for selected rig.
func (m *miningPrivate) GetRigStats() {
	//TODO: Implement method
}

//GetRigStatisticsByAlgo get statistical streams for selected rig and selected algorithm.
func (m *miningPrivate) GetRigStatisticsByAlgo(algo string) {
	//TODO: Implement method
}

//GetRigDetails get miningPrivate rig detailed information for selected rig.
func (m *miningPrivate) GetRigDetails(rigId string) (*rig, error) {
	m.client.doRequest()

	/*resp, err := m.httpClient.R().
		SetResult(&rig{}).
		Get(fmt.Sprintf("%s/main/api/v2/miningPrivate/rig2/%s", ProdUrl, rigId))

	if err != nil {
		return nil, errors.Wrap(err, "unexpected error")
	}

	if requestErr, ok := resp.Error().(*requestError); ok {
		return nil, errors.New(requestErr.Errors[0].Message)
	}

	rig, ok := resp.Result().(*rig)
	if !ok {
		return nil, errors.New("invalid result")
	}

	return rig, nil*/

	return nil, nil
}

//GetActiveWorkers get a list of active worker.
func (m *miningPrivate) GetActiveWorkers() {
	//TODO: Implement method
}

//GetPayouts get list of payouts.
func (m *miningPrivate) GetPayouts() {
	//TODO: Implement method
}

// GetAllRigsStatistics get statistical streams for all miningPrivate rigs.
func (m *miningPrivate) GetAllRigsStatistics() {
	//TODO: Implement method
}

//GetAllRigsStatisticsByAlgo get statistical streams for all miningPrivate rigs for selected algorithm. Algorithm code can be found in buy info endpoint.
func (m *miningPrivate) GetAllRigsStatisticsByAlgo(algo string) {
	//TODO: Implement method
}

//GetAllRigs list rigs and their statuses.
//Path parameter filters rigs by group.
//When path is empty, rigs from root group are returned.
//Rigs can be sorted according to sort parameter.
func (m *miningPrivate) GetAllRigs(algo string) {
	//TODO: Implement method
}

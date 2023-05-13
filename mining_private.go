package nhclient

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type optionalParameter struct {
	key, value string
}

type miningPrivate struct {
	client *client
}

// GetAddress get the mining address.
func (m *miningPrivate) GetAddress() (string, error) {
	responseBody, err := m.client.doRequest(
		"GET",
		"/main/api/v2/mining/miningAddress",
		nil,
		nil,
	)

	if err != nil {
		return "", errors.Wrap(err, "executing request")
	}

	var response struct {
		Address string
	}
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return "", errors.Wrap(err, "unmarshalling response body")
	}

	return response.Address, nil
}

// TODO: Implement struct
type rigsGroups struct{}

// GetRigsGroups list groups with list of rigs in the groups.
// When extendedResponse is set to true, response contains number of total and active devices for each rig and group.
func (m *miningPrivate) GetRigsGroups(parameters ...*optionalParameter) (rg *rigsGroups, err error) {
	//TODO: Optional query parameter extendedResponse (bool)
	path := "/main/api/v2/mining/groups/list"

	queryParameters := make(map[string]string, 0)
	for _, parameter := range parameters {
		queryParameters[parameter.key] = parameter.value
	}

	responseBody, err := m.client.doRequest(
		"GET",
		path,
		nil,
		queryParameters,
	)

	if err != nil {
		return nil, errors.Wrap(err, "executing request")
	}

	err = json.Unmarshal(responseBody, &rg)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling response body")
	}

	return rg, nil
}

// TODO: Implement struct
type algoStatistics struct{}

// GetAlgoStatistics list mining algos with basic statistics for organization (and for rig id if specified).
func (m *miningPrivate) GetAlgoStatistics() (algoStats *algoStatistics, err error) {
	//TODO: Optional query parameter rigId (string)
	path := "/main/api/v2/mining/algo/stats"

	responseBody, err := m.client.doRequest(
		"GET",
		path,
		nil,
		nil,
	)

	if err != nil {
		return nil, errors.Wrap(err, "executing request")
	}

	err = json.Unmarshal(responseBody, &algoStats)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling response body")
	}

	return algoStats, nil
}

// GetRigStatistics get statistical streams for selected rig.
func (m *miningPrivate) GetRigStatisticsByAlgo() error {
	//TODO: Implement method
	return errNotImplemented
}

// GetRigUnpaidStatistics get statistical streams for unpaid amounts
func (m *miningPrivate) GetRigUnpaidStatistics() error {
	//TODO: Implement method
	return errNotImplemented
}

// GetMinerStatisticsByAlgo get statistical streams for all mining rigs by algorithm.
func (m *miningPrivate) GetMinerStatisticsByAlgo() error {
	//TODO: Implement method
	return errNotImplemented
}

// GetAllRigsStatisticsByAlgo get statistical streams for all mining rigs for selected algorithm. Algorithm code can be found in buy info endpoint.
func (m *miningPrivate) GetMinerUnpaidStatistics() error {
	//TODO: Implement method
	return errNotImplemented
}

// GetRigDetails get mining rig detailed information for selected rig.
func (m *miningPrivate) GetRigDetails(rigId string) (rig *rig, err error) {
	path := fmt.Sprintf("/main/api/v2/mining/rig2/%s", rigId)

	responseBody, err := m.client.doRequest("GET",
		path,
		nil,
		nil,
	)

	if err != nil {
		return nil, errors.Wrap(err, "executing request")
	}

	err = json.Unmarshal(responseBody, &rig)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling response body")
	}

	return rig, nil
}

// GetActiveWorkers get a list of active worker.
func (m *miningPrivate) GetActiveWorkers() error {
	//TODO: Implement method
	return errNotImplemented
}

// GetPayouts get list of payouts.
func (m *miningPrivate) GetPayouts() error {
	//TODO: Implement method
	return errNotImplemented
}

// GetAllRigs list rigs and their statuses.
// Path parameter filters rigs by group.
// When path is empty, rigs from root group are returned.
// Rigs can be sorted according to sort parameter.
func (m *miningPrivate) GetAllRigs(algo string) error {
	//TODO: Implement method
	return errNotImplemented
}

//------------------------//
//TODO: Implement POST method

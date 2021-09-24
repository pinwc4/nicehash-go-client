package nhclient

type mining struct {}

//GetAddress get the mining address.
func (m *mining) GetAddress() {
	//TODO: Implement method
}

//GetAlgos list mining algos with basic statistics for organization (and for rig id if specified).
func (m *mining) GetAlgos(rigId *string)  {
	//TODO: Implement method
}

//GetRigsGroups list groups with list of rigs in the groups.
//When extendedResponse is set to true, response contains number of total and active devices for each rig and group.
func (m *mining) GetRigsGroups()  {
	//TODO: Implement method
}

//GetRigStats get statistical streams for selected rig.
func (m *mining) GetRigStats()  {
	//TODO: Implement method
}

//GetRigStatisticsByAlgo get statistical streams for selected rig and selected algorithm.
func (m *mining) GetRigStatisticsByAlgo(algo string)  {
	//TODO: Implement method
}

//GetRigDetails get mining rig detailed information for selected rig.
func (m *mining) GetRigDetails(rigId string)  {
	//TODO: Implement method
}

//GetActiveWorkers get a list of active worker.
func (m *mining) GetActiveWorkers()  {
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



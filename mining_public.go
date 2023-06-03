package nhclient

type miningPublic struct {
	client *Client
}

func (m *miningPublic) GetActiveWorkersByBTCAddress() error {
	return errNotImplemented
}

func (m *miningPublic) GetMinerStatisticsByBTCAddressAndAlgo() error {
	return errNotImplemented
}
func (m *miningPublic) GetMinerStatisticsByBTCAddress() error {
	return errNotImplemented
}
func (m *miningPublic) GetWithdrawalsByBTCAddress() error {
	return errNotImplemented
}
func (m *miningPublic) GetMinerStatusByBTCAddress() error {
	return errNotImplemented
}

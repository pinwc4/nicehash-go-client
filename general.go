package nhclient

type general struct {
	client *client
}

func (g *general) GetMiningAlgorithms() error {
	g.client.doRequest(
		"GET",
		"/main/api/v2/mining/algorithms",
		nil,
		nil,
	)

	return errNotImplemented
}

func (g *general) GetMiningMarkets() error {
	return errNotImplemented
}

func (g *general) GetCurrencies() error {
	return errNotImplemented
}
func (g *general) GetFeeRules() error {
	return errNotImplemented
}
func (g *general) GetCountries() error {
	return errNotImplemented
}
func (g *general) GetAllowedKMCountries() error {
	return errNotImplemented
}
func (g *general) GetPossiblePermissions() error {
	return errNotImplemented
}
func (g *general) GetAllowedXCHCountries() error {
	return errNotImplemented
}
func (g *general) GetAPIFlags() error {
	return errNotImplemented
}
func (g *general) GetServerTime() error {
	return errNotImplemented
}

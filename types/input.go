package types

type ItemInput struct {
	Quantity int    `json:"quantity"`
	ID       string `json:"id"`
}
type PriceSimulationInput struct {
	Items []ItemInput `json:"items"`
}

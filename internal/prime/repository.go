package prime

import (
	"context"

	"prime/pkg/vtex"
)

type PrimeRepository interface {
	SimulationPrice(ctx context.Context, input vtex.ReqSearchProducts) (*vtex.ResSearchProducts, error)
	PrimeConfig(ctx context.Context) (*vtex.ResPrimeConfig, error)
}

type primeRepository struct {
	client *vtex.Client
}

func NewPrimeRepository(client *vtex.Client) PrimeRepository {
	return &primeRepository{client: client}
}

func (p *primeRepository) SimulationPrice(ctx context.Context, input vtex.ReqSearchProducts) (*vtex.ResSearchProducts, error) {
	data, err := p.client.SearchProductsBySkuIds(ctx, input)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *primeRepository) PrimeConfig(ctx context.Context) (*vtex.ResPrimeConfig, error) {
	data, err := p.client.GetPrimeConfig(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

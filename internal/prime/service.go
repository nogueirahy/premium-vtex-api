package prime

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"prime/pkg/cache"
	"prime/pkg/vtex"
	"prime/types"
)

type PrimeService interface {
	SimulationPrice(ctx context.Context, input types.PriceSimulationInput) (PriceSimulationOutput, error)
}

type primeService struct {
	primeRepo PrimeRepository
}

func NewPrimeService(primeRepo PrimeRepository) PrimeService {
	return &primeService{primeRepo: primeRepo}
}

const (
	PRIME_KEY        = "PRIME"
	PRIME_KEY_PREFIX = "PRIME_PRODUCT_"
)

func buildQuery(items []types.ItemInput) string {
	var parts []string
	for _, item := range items {
		parts = append(parts, "fq=skuId:"+item.ID)
	}
	return strings.Join(parts, "&")
}

func (p *primeService) SimulationPrice(ctx context.Context, input types.PriceSimulationInput) (PriceSimulationOutput, error) {
	cache, err := cache.GetCache()
	if err != nil {
		return PriceSimulationOutput{}, err
	}

	var products vtex.ResSearchProducts
	var missingProducts []types.ItemInput

	for _, item := range input.Items {
		cachedProduct, found := cache.Get(PRIME_KEY_PREFIX + item.ID)

		if found {
			var product vtex.Product
			err := json.Unmarshal([]byte(cachedProduct), &product)
			if err != nil {
				return PriceSimulationOutput{}, fmt.Errorf("failed to unmarshal product from cache: %w", err)
			}
			products = append(products, product)
		} else {
			missingProducts = append(missingProducts, item)
		}
	}

	if len(missingProducts) > 0 {
		query := buildQuery(missingProducts)

		newProducts, err := p.primeRepo.SimulationPrice(ctx, vtex.ReqSearchProducts{Query: query})
		if err != nil {
			return PriceSimulationOutput{}, err
		}

		products = append(products, *newProducts...)

		for _, product := range *newProducts {
			for _, item := range product.Items {
				jsonProduct, err := json.Marshal(product)
				if err != nil {
					return PriceSimulationOutput{}, err
				}
				cache.Set(PRIME_KEY_PREFIX+item.ItemID, string(jsonProduct), time.Hour)
			}
		}
	}

	var primeConfig *vtex.ResPrimeConfig

	cachedPrimeConfig, found := cache.Get(PRIME_KEY)

	if !found {
		primeConfig, _ := p.primeRepo.PrimeConfig(ctx)
		cachedPrimeConfig, err := json.Marshal(primeConfig)
		if err != nil {
			return PriceSimulationOutput{}, err
		}
		cache.Set(PRIME_KEY, string(cachedPrimeConfig), time.Hour)
	} else {
		err = json.Unmarshal([]byte(cachedPrimeConfig), &primeConfig)
		if err != nil {
			return PriceSimulationOutput{}, err
		}
	}

	price := adaptPricePrime(input, products, primeConfig)

	return PriceSimulationOutput{
		Price: price,
	}, nil
}

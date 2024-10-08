package prime

import (
	"strings"

	"prime/pkg/vtex"
	"prime/types"
)

func getDefaultSeller(sellers []vtex.Seller) *string {
	for _, seller := range sellers {
		if seller.SellerDefault {
			return &seller.SellerID
		}
	}
	if len(sellers) > 0 {
		return &sellers[0].SellerID
	}
	return nil
}

func hasPrimeConditions(product vtex.Product, promo vtex.ResPrimeConfig) bool {
	if len(product.Items) == 0 || len(product.Items[0].Sellers) == 0 {
		return false
	}

	commertialOffer := product.Items[0].Sellers[0].CommertialOffer
	if commertialOffer.ListPrice != commertialOffer.Price || promo.PercentualDiscountValue == 0 {
		return false
	}

	categories := strings.Join(mapCategories(promo.Categories), ",")
	if promo.CategoriesAreInclusive != strings.Contains(categories, product.CategoryID) {
		return false
	}

	brands := strings.Join(mapBrands(promo.Brands), ",")
	if promo.BrandsAreInclusive != strings.Contains(brands, string(rune(product.BrandID))) {
		return false
	}

	collections := mapCollections(promo.Collections)
	productCollections := mapProductClusters(product.ProductClusters)
	if promo.CollectionsIsInclusive != containsAny(collections, productCollections) {
		return false
	}

	sellers := getProductSellers(product.Items)

	return promo.IDSellerIsInclusive == containsAnyString(promo.IDSeller, sellers)
}

func mapCategories(categories []vtex.Category) []string {
	mapped := make([]string, len(categories))
	for i, c := range categories {
		mapped[i] = c.ID
	}
	return mapped
}

func mapBrands(brands []vtex.Brand) []string {
	mapped := make([]string, len(brands))
	for i, b := range brands {
		mapped[i] = b.ID
	}
	return mapped
}

func mapCollections(collections []vtex.Collection) []string {
	mapped := make([]string, len(collections))
	for i, c := range collections {
		mapped[i] = c.ID
	}
	return mapped
}

func mapProductClusters(clusters map[string]any) []string {
	mapped := make([]string, 0, len(clusters))
	for k := range clusters {
		mapped = append(mapped, k)
	}
	return mapped
}

func containsAny(slice1, slice2 []string) bool {
	set := make(map[string]struct{}, len(slice1))
	for _, item := range slice1 {
		set[item] = struct{}{}
	}
	for _, item := range slice2 {
		if _, found := set[item]; found {
			return true
		}
	}
	return false
}

func getProductSellers(items []vtex.Item) []string {
	sellers := make([]string, 0, len(items))
	for _, item := range items {
		seller := getDefaultSeller(item.Sellers)
		if seller != nil {
			sellers = append(sellers, *seller)
		}
	}
	return sellers
}

func containsAnyString(str string, slice []string) bool {
	for _, s := range slice {
		if strings.Contains(str, s) {
			return true
		}
	}
	return false
}

func adaptPricePrime(input types.PriceSimulationInput, data vtex.ResSearchProducts, primeConfig *vtex.ResPrimeConfig) float64 {
	var discountPrime, primePrice float64
	var quantity int = 1

	for _, product := range data {

		for _, itemProduct := range product.Items {
			for _, item := range input.Items {
				if itemProduct.ItemID == item.ID {
					quantity = item.Quantity
					break
				}
			}
		}

		commertialOffer := product.Items[0].Sellers[0].CommertialOffer
		if commertialOffer.ListPrice != commertialOffer.Price || primeConfig.PercentualDiscountValue == 0 {
			discountPrime += commertialOffer.Price
		} else if hasPrimeConditions(product, *primeConfig) {
			priceWithDiscount := commertialOffer.Price
			primePercentualDiscount := primeConfig.PercentualDiscountValue
			primePrice = priceWithDiscount * (1 - float64(primePercentualDiscount)/100)
			discountPrime += priceWithDiscount - primePrice
		}
	}

	return discountPrime * float64(quantity)
}

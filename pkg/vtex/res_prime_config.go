package vtex

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Collection struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Brand struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ResPrimeConfig struct {
	Name                      string       `json:"name"`
	IsActive                  bool         `json:"isActive"`
	IDCalculatorConfiguration string       `json:"idCalculatorConfiguration"`
	PercentualDiscountValue   int          `json:"percentualDiscountValue"`
	MarketingTags             []string     `json:"marketingTags"`
	IDSeller                  string       `json:"idSeller"`
	Categories                []Category   `json:"categories"`
	Brands                    []Brand      `json:"brands"`
	Collections               []Collection `json:"collections"`
	CategoriesAreInclusive    bool         `json:"categoriesAreInclusive"`
	IDSellerIsInclusive       bool         `json:"idSellerIsInclusive"`
	BrandsAreInclusive        bool         `json:"brandsAreInclusive"`
	CollectionsIsInclusive    bool         `json:"collectionsIsInclusive"`
}

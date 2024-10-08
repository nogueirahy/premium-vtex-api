package vtex

import "time"

type ResSearchProducts []Product

type Product struct {
	Items                   []Item             `json:"items"`
	ProductID               string             `json:"productId"`
	ProductName             string             `json:"productName"`
	Brand                   string             `json:"brand"`
	BrandID                 int                `json:"brandId"`
	BrandImageURL           interface{}        `json:"brandImageUrl"`
	LinkText                string             `json:"linkText"`
	ProductReference        string             `json:"productReference"`
	ProductReferenceCode    string             `json:"productReferenceCode"`
	CategoryID              string             `json:"categoryId"`
	ProductTitle            string             `json:"productTitle"`
	MetaTagDescription      interface{}        `json:"metaTagDescription"`
	ReleaseDate             time.Time          `json:"releaseDate"`
	ClusterHighlights       map[string]any     `json:"clusterHighlights"`
	ProductClusters         map[string]any     `json:"productClusters"`
	SearchableClusters      map[string]any     `json:"searchableClusters"`
	Categories              []string           `json:"categories"`
	CategoriesIds           []string           `json:"categoriesIds"`
	Link                    string             `json:"link"`
	ComposiO                []string           `json:"Composição"`
	EspecificaEs            []string           `json:"Especificações"`
	Gender                  []string           `json:"gender"`
	AgeGroup                []string           `json:"age_group"`
	AtributosGoogle         []string           `json:"Atributos Google"`
	AllSpecifications       []string           `json:"allSpecifications"`
	AllSpecificationsGroups []string           `json:"allSpecificationsGroups"`
	Description             string             `json:"description"`
	SkuSpecifications       []SkuSpecification `json:"skuSpecifications"`
}

type Item struct {
	Sellers              []Seller      `json:"sellers"`
	UnitMultiplier       float64       `json:"unitMultiplier"`
	ItemID               string        `json:"itemId"`
	Name                 string        `json:"name"`
	NameComplete         string        `json:"nameComplete"`
	ComplementName       string        `json:"complementName"`
	Ean                  string        `json:"ean"`
	ReferenceID          []ReferenceID `json:"referenceId"`
	MeasurementUnit      string        `json:"measurementUnit"`
	ModalType            interface{}   `json:"modalType"`
	IsKit                bool          `json:"isKit"`
	Images               []Image       `json:"images"`
	Tamanho              []string      `json:"Tamanho"`
	IDCORORIGINAL        []string      `json:"ID_COR_ORIGINAL"`
	VALORHEXORIGINAL     []string      `json:"VALOR_HEX_ORIGINAL"`
	DESCCORORIGINAL      []string      `json:"DESC_COR_ORIGINAL"`
	IDCORCONSOLIDADA     []string      `json:"ID_COR_CONSOLIDADA"`
	VALORHEXCONSOLIDADA  []string      `json:"VALOR_HEX_CONSOLIDADA"`
	COR                  []string      `json:"COR"`
	Variations           []string      `json:"variations"`
	Attachments          []Attachment  `json:"attachments"`
	Videos               []interface{} `json:"Videos"`
	EstimatedDateArrival interface{}   `json:"estimatedDateArrival"`
}

type ReferenceID struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type Image struct {
	ImageID           string      `json:"imageId"`
	ImageLabel        interface{} `json:"imageLabel"`
	ImageTag          string      `json:"imageTag"`
	ImageURL          string      `json:"imageUrl"`
	ImageText         interface{} `json:"imageText"`
	ImageLastModified time.Time   `json:"imageLastModified"`
}

type Attachment struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Required     bool   `json:"required"`
	DomainValues string `json:"domainValues"`
}

type Seller struct {
	CommertialOffer CommertialOffer `json:"commertialOffer"`
	SellerID        string          `json:"sellerId"`
	SellerName      string          `json:"sellerName"`
	AddToCartLink   string          `json:"addToCartLink"`
	SellerDefault   bool            `json:"sellerDefault"`
}

type CommertialOffer struct {
	Price                          float64                       `json:"Price"`
	ListPrice                      float64                       `json:"ListPrice"`
	PriceWithoutDiscount           float64                       `json:"PriceWithoutDiscount"`
	FullSellingPrice               float64                       `json:"FullSellingPrice"`
	RewardValue                    float64                       `json:"RewardValue"`
	Tax                            float64                       `json:"Tax"`
	DeliverySLASamplesPerRegion    DeliverySLASamplesPerRegion   `json:"DeliverySlaSamplesPerRegion"`
	Installments                   []Installment                 `json:"Installments"`
	DiscountHighLight              []interface{}                 `json:"DiscountHighLight"`
	GiftSkuIds                     []interface{}                 `json:"GiftSkuIds"`
	Teasers                        []interface{}                 `json:"Teasers"`
	PromotionTeasers               []interface{}                 `json:"PromotionTeasers"`
	BuyTogether                    []interface{}                 `json:"BuyTogether"`
	ItemMetadataAttachment         []interface{}                 `json:"ItemMetadataAttachment"`
	PriceValidUntil                time.Time                     `json:"PriceValidUntil"`
	AvailableQuantity              int                           `json:"AvailableQuantity"`
	IsAvailable                    bool                          `json:"IsAvailable"`
	DeliverySLASamples             []DeliverySLASamplesPerRegion `json:"DeliverySlaSamples"`
	GetInfoErrorMessage            interface{}                   `json:"GetInfoErrorMessage"`
	CacheVersionUsedToCallCheckout string                        `json:"CacheVersionUsedToCallCheckout"`
	PaymentOptions                 PaymentOptions                `json:"PaymentOptions"`
}

type DeliverySLASamplesPerRegion struct {
	Num0 DeliverySLA `json:"0"`
}

type DeliverySLA struct {
	DeliverySLAPerTypes []interface{} `json:"DeliverySlaPerTypes"`
	Region              interface{}   `json:"Region"`
}

type Installment struct {
	Value                      float64 `json:"Value"`
	InterestRate               float64 `json:"InterestRate"`
	TotalValuePlusInterestRate float64 `json:"TotalValuePlusInterestRate"`
	NumberOfInstallments       int     `json:"NumberOfInstallments"`
	PaymentSystemName          string  `json:"PaymentSystemName"`
	PaymentSystemGroupName     string  `json:"PaymentSystemGroupName"`
	Name                       string  `json:"Name"`
}

type PaymentOptions struct {
	InstallmentOptions []InstallmentOption `json:"installmentOptions"`
	PaymentSystems     []PaymentSystem     `json:"paymentSystems"`
	Payments           []interface{}       `json:"payments"`
	GiftCards          []interface{}       `json:"giftCards"`
	GiftCardMessages   []interface{}       `json:"giftCardMessages"`
	AvailableAccounts  []interface{}       `json:"availableAccounts"`
	AvailableTokens    []interface{}       `json:"availableTokens"`
}

type InstallmentOption struct {
	PaymentSystem    string        `json:"paymentSystem"`
	Bin              interface{}   `json:"bin"`
	PaymentName      string        `json:"paymentName"`
	PaymentGroupName string        `json:"paymentGroupName"`
	Value            int           `json:"value"`
	Installments     []Installment `json:"installments"`
}

type PaymentSystem struct {
	ID                     int         `json:"id"`
	Name                   string      `json:"name"`
	GroupName              string      `json:"groupName"`
	Validator              interface{} `json:"validator"`
	StringID               string      `json:"stringId"`
	Template               string      `json:"template"`
	RequiresDocument       bool        `json:"requiresDocument"`
	IsCustom               bool        `json:"isCustom"`
	Description            interface{} `json:"description"`
	RequiresAuthentication bool        `json:"requiresAuthentication"`
	DueDate                time.Time   `json:"dueDate"`
	AvailablePayments      interface{} `json:"availablePayments"`
}

type SkuSpecification struct {
	Field  SkuField   `json:"field"`
	Values []SkuValue `json:"values"`
}

type SkuField struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
	Position int    `json:"position"`
	Type     string `json:"type"`
}

type SkuValue struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
}

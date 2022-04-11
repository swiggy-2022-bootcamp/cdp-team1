package entity

type Product struct {
	ID                 ID                   `json:"id"`
	Model              string               `json:"model"`
	Quantity           string               `json:"quantity"`
	Price              string               `json:"price"`
	ManufacturerId     string               `json:"manufacturer_id"`
	Sku                string               `json:"sku"`
	ProductSeoUrl      []ProductSeoUrl      `json:"product_seo_url"`
	Points             int                  `json:"points"`
	Reward             int                  `json:"reward"`
	Image              string               `json:"image"`
	Shipping           string               `json:"shipping"`
	Weight             int                  `json:"weight"`
	Length             int                  `json:"length"`
	Width              int                  `json:"width"`
	Height             int                  `json:"height"`
	Minimum            int                  `json:"minimum"`
	ProductRelated     []string             `json:"product_related"`
	ProductDescription []ProductDescription `json:"product_description"`
	ProductCategory    []string             `json:"product_category"`
}

type ProductSeoUrl struct {
	Keyword    string `json:"keyword"`
	LanguageId string `json:"language_id"`
	StoreId    string `json:"store_id"`
}

type ProductDescription struct {
	LanguageId      string `json:"language_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	MetaKeyword     string `json:"meta_keyword"`
	Tag             string `json:"tag"`
}

func NewProduct(model, quantity, price, manufacturerId, sku string,
	productSeoUrl []ProductSeoUrl,
	points, reward int,
	image, shipping string,
	weight, length, width, height, minimum int,
	productRelated []string,
	productDescription []ProductDescription,
	productCategory []string) (*Product, error) {

	product := &Product{
		ID:                 NewID(),
		Model:              model,
		Quantity:           quantity,
		Price:              price,
		ManufacturerId:     manufacturerId,
		Sku:                sku,
		ProductSeoUrl:      productSeoUrl,
		Points:             points,
		Reward:             reward,
		Image:              image,
		Shipping:           shipping,
		Weight:             weight,
		Length:             length,
		Width:              width,
		Height:             height,
		Minimum:            minimum,
		ProductRelated:     productRelated,
		ProductDescription: productDescription,
		ProductCategory:    productCategory,
	}

	err := product.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.Model == "" || p.ManufacturerId == "" || p.Sku == "" || p.ProductSeoUrl == nil || p.Points < 0 ||
		p.Reward < 0 || p.Weight < 0 || p.Height < 0 || p.Length < 0 || p.Minimum < 0 || p.ProductDescription == nil ||
		p.ProductCategory == nil {
		return ErrInvalidEntity
	}
	return nil
}

func NewProductSeoUrl(keyWord, languageId, storeId string) ProductSeoUrl {
	productSeoUrl := ProductSeoUrl{
		Keyword:    keyWord,
		LanguageId: languageId,
		StoreId:    storeId,
	}
	return productSeoUrl
}

func NewProductDescription(languageId, name, description, metaTitle, metaDescription, metaKeyword, tag string) ProductDescription {
	productDescription := ProductDescription{
		LanguageId:      languageId,
		Name:            name,
		Description:     description,
		MetaTitle:       metaTitle,
		MetaDescription: metaDescription,
		MetaKeyword:     metaKeyword,
		Tag:             tag,
	}
	return productDescription
}

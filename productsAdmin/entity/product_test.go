package entity_test

import (
	"github.com/stretchr/testify/assert"
	"qwik.in/productsAdmin/entity"
	"testing"
)

func TestNewProductSeoUrl(t *testing.T) {
	p := entity.NewProductSeoUrl("demo-keyword", "1", "0")
	assert.Equal(t, p.Keyword, "demo-keyword")
	assert.Equal(t, p.LanguageId, "1")
	assert.Equal(t, p.StoreId, "0")
}

func TestNewProductDescription(t *testing.T) {
	p := entity.NewProductDescription(
		"1",
		"demo name",
		"Description of the product",
		"Meta title of the product",
		"Meta description of the product",
		"demo, keyword, demo2",
		"Product's tag, comma separated",
	)
	assert.Equal(t, p.LanguageId, "1")
	assert.Equal(t, p.Name, "demo name")
	assert.Equal(t, p.Description, "Description of the product")
	assert.Equal(t, p.MetaTitle, "Meta title of the product")
	assert.Equal(t, p.MetaDescription, "Meta description of the product")
	assert.Equal(t, p.MetaKeyword, "demo, keyword, demo2")
	assert.Equal(t, p.Tag, "Product's tag, comma separated")
}

func TestNewProduct(t *testing.T) {

	p, err := entity.NewProduct(
		"demo model",
		"300",
		"44.00",
		"20",
		"demo sku",
		[]entity.ProductSeoUrl{entity.NewProductSeoUrl("demo-keyword", "1", "0")},
		0,
		0,
		"",
		"1",
		0,
		0,
		0,
		0,
		0,
		[]string{"34"},
		[]entity.ProductDescription{entity.NewProductDescription(
			"1",
			"demo name",
			"Description of the product",
			"Meta title of the product",
			"Meta description of the product",
			"demo, keyword, demo2",
			"Product's tag, comma separated",
		)},
		[]string{"25"},
	)
	assert.Nil(t, err)
	assert.NotNil(t, p.ID)
	assert.Equal(t, p.Model, "demo model")
	assert.Equal(t, p.Quantity, "300")
	assert.Equal(t, p.Price, "44.00")
	assert.Equal(t, p.ManufacturerId, "20")
	assert.Equal(t, p.Sku, "demo sku")
	assert.NotNil(t, p.ProductSeoUrl)
	assert.Equal(t, p.Points, 0)
	assert.Equal(t, p.Reward, 0)
	assert.Equal(t, p.Image, "")
	assert.Equal(t, p.Shipping, "1")
	assert.Equal(t, p.Weight, 0)
	assert.Equal(t, p.Length, 0)
	assert.Equal(t, p.Width, 0)
	assert.Equal(t, p.Height, 0)
	assert.Equal(t, p.Minimum, 0)
	assert.NotNil(t, p.ProductRelated)
	assert.NotNil(t, p.ProductDescription)
	assert.NotNil(t, p.ProductCategory)

}

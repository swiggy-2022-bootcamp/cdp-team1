package entity_test

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"os"
	"qwik.in/productsAdmin/entity"
	"qwik.in/productsAdmin/log"
	"reflect"
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
	assert.Equal(t, len(p.ProductSeoUrl), 1)
	assert.Equal(t, reflect.DeepEqual(p.ProductSeoUrl[0], entity.NewProductSeoUrl("demo-keyword", "1", "0")), true)
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
	assert.Equal(t, len(p.ProductRelated), 1)
	assert.Equal(t, p.ProductRelated[0], "34")
	assert.NotNil(t, p.ProductDescription)
	assert.Equal(t, len(p.ProductDescription), 1)
	assert.Equal(t, reflect.DeepEqual(p.ProductDescription[0], entity.NewProductDescription(
		"1",
		"demo name",
		"Description of the product",
		"Meta title of the product",
		"Meta description of the product",
		"demo, keyword, demo2",
		"Product's tag, comma separated",
	)), true)
	assert.NotNil(t, p.ProductCategory)
	assert.Equal(t, len(p.ProductCategory), 1)
	assert.Equal(t, p.ProductCategory[0], "25")
}

func TestSetId(t *testing.T) {

	productData, _ := os.ReadFile("sampleProductSingle.json")
	var p entity.Product
	if err := json.Unmarshal(productData, &p); err != nil {
		log.Error(err)
	}

	assert.Equal(t, "90f97384-0117-4234-bbe2-5f306bbef0b3", p.ID)
	p.SetId()
	assert.NotEqual(t, "", p.ID)
}

func TestProductValidate(t *testing.T) {
	_, err := entity.NewProduct(
		"",
		"",
		"",
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

	assert.Equal(t, err.Error(), "Invalid entity")
}

func TestReadFromFile(t *testing.T) {

	productData, _ := os.ReadFile("sampleProducts.json")
	var products []entity.Product
	if err := json.Unmarshal(productData, &products); err != nil {
		log.Error(err)
	}

	assert.Equal(t, 3, len(products))

	productData1, _ := os.ReadFile("sampleProductSingle.json")
	var product1 entity.Product
	if err := json.Unmarshal(productData1, &product1); err != nil {
		log.Error(err)
	}

	assert.True(t, cmp.Equal(products[0], product1))
}

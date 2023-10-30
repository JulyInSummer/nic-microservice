package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Products []*Product

var ErrProductNotFound = fmt.Errorf("no product matching this ID")

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// id for this product
	//
	// required: true
	// min: 1
	ID          int     `json:"id"`

	// the name for this product
	//
	// required: true
	// max length: 255
	Name        string  `json:"name" validate:"required"`

	// the description for this product
	//
	// required: true
	// max length: 1000
	Description string  `json:"description"`

	// the price for this product
	//
	// required: true
	// min: 0.1
	Price       float32 `json:"price" validate:"gt=0"`

	// the SKU for this product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// A list of products return in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All the products in the system
	// in: body
	Body []Product
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func GetProducts() Products {
	return productList
}

func UpdateProduct(id int, product *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	product.ID = id
	productList[pos] = product
	return nil
}

func DeleteProduct(id int) error {
	for pos, product := range productList {
		if product.ID == id {
			productList = append(productList[:pos], productList[pos+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Couldn't delete product with id:", id)
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	nextID := lp.ID + 1
	return nextID
}

func validateSKU(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := reg.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

var productList = Products{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

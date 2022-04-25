package service

import (
	"cartService/domain/model"
	"cartService/domain/repository"
	"cartService/internal/error"
	"fmt"
)

type CartService interface {
	AddToCart(*model.Cart, string) *error.AppError
	GetAllCart() (*[]model.Cart, *error.AppError)
	UpdateCart(string, string, int) *error.AppError
	DeleteCartByCustomerId(string) *error.AppError
	DeleteAllCart() *error.AppError
}

type CartServiceImpl struct {
	cartRepository repository.CartRepositoryDB
}

func NewCartService(cartRepository repository.CartRepositoryDB) CartService {
	return &CartServiceImpl{
		cartRepository: cartRepository,
	}
}

func (csvc CartServiceImpl) AddToCart(cart *model.Cart, customer_id string) *error.AppError {

	// Check if cart already exists
	// If exists, update the quantity
	// If not, create a new cart
	curr_cart, err := csvc.cartRepository.Read(customer_id)

	fmt.Println(curr_cart, err)

	if err == nil {

		// curr_cart.Products = append(curr_cart.Products, *product)

		updated_cart := model.Cart{
			CustomerId: customer_id,
			Products:   curr_cart.Products,
		}

		err2 := csvc.cartRepository.UpdateExisting(&updated_cart)

		if err2 != nil {
			return err2
		}

		return nil
	}

	fmt.Println("cart doesnt exist")
	fmt.Println("Product: ", cart)

	// new_cart := model.Cart{
	// 	CustomerId: customer_id,
	// 	Products:   []model.C{*product},
	// }

	// fmt.Println("new cart", new_cart)

	// err3 := csvc.cartRepository.Create(new_cart)

	// if err3 != nil {
	// return err3
	// }

	return nil
}

func (csvc CartServiceImpl) GetAllCart() (*[]model.Cart, *error.AppError) {

	u, err := csvc.cartRepository.ReadAll()

	if err != nil {
		return nil, err
	}

	return u, err
}

func (csvc CartServiceImpl) UpdateCart(customer_id string, product_id string, quantity int) *error.AppError {

	// Get current cart object
	curr_cart, err := csvc.cartRepository.Read(customer_id)

	if err != nil {
		return err
	}

	// Find the product in the cart
	// Update the quantity
	for i, p := range curr_cart.Products {
		if p.ProductId == product_id {
			curr_cart.Products[i].Quantity = quantity
		}
	}

	err2 := csvc.cartRepository.UpdateExisting(curr_cart)

	if err2 != nil {
		return err
	}

	return err2
}

func (csvc CartServiceImpl) DeleteCartByCustomerId(customer_id string) *error.AppError {

	err := csvc.cartRepository.Delete(customer_id)

	if err != nil {
		return err
	}

	return err
}

func (csvc CartServiceImpl) DeleteAllCart() *error.AppError {

	err := csvc.cartRepository.DeleteAll()

	if err != nil {
		return err
	}

	return nil
}

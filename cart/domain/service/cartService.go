package service

import (
	"cartService/db"
	"cartService/domain/model"
	"cartService/domain/repository"
	app_error "cartService/internal/error"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type CartService interface {
	AddToCart(*model.Cart, string) *app_error.AppError
	GetCartByCustomerId(string) (*model.Cart, *app_error.AppError)
	UpdateCart(string, string, int) *app_error.AppError
	DeleteCartByCustomerId(string) *app_error.AppError
	DeleteCartItem(string, string) *app_error.AppError
	UserIDFromAuthToken(string) (string, *app_error.AppError)
}

type CartServiceImpl struct {
	cartRepository repository.CartRepositoryDB
}

func NewCartService(cartRepository repository.CartRepositoryDB) CartService {
	return &CartServiceImpl{
		cartRepository: cartRepository,
	}
}

func (csvc CartServiceImpl) AddToCart(cart *model.Cart, customer_id string) *app_error.AppError {

	// Check if cart already exists
	// If exists, update the quantity
	// If not, create a new cart
	curr_cart, err := csvc.cartRepository.Read(customer_id)

	fmt.Println("curr cart: ", curr_cart, err)

	if err == nil {

		curr_cart.Products = append(curr_cart.Products, cart.Products...)

		fmt.Println("curr cart updated: ", curr_cart)

		err2 := csvc.cartRepository.UpdateExisting(curr_cart)

		if err2 != nil {
			return err2
		}

		return nil
	}

	fmt.Println("cart doesnt exist")
	fmt.Println("Cart: ", cart)
	fmt.Println("Products: ", cart.Products)
	fmt.Println("inside products: ", cart.Products[0].ProductId)
	fmt.Println("inside quantity: ", cart.Products[0].Quantity)

	new_product := model.Product{
		ProductId: cart.Products[0].ProductId,
		Quantity:  cart.Products[0].Quantity,
	}

	fmt.Println("new product: ", new_product)

	new_cart := model.Cart{
		CustomerId: customer_id,
		Products:   []model.Product{new_product},
	}

	fmt.Println("new cart: ", new_cart)

	err3 := csvc.cartRepository.Create(new_cart)
	if err3 != nil {
		return err3
	}

	return nil
}

func (csvc CartServiceImpl) GetCartByCustomerId(customer_id string) (*model.Cart, *app_error.AppError) {

	u, err := csvc.cartRepository.Read(customer_id)

	if err != nil {
		return nil, err
	}

	return u, err
}

func (csvc CartServiceImpl) UpdateCart(customer_id string, product_id string, quantity int) *app_error.AppError {

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

func (csvc CartServiceImpl) DeleteCartByCustomerId(customer_id string) *app_error.AppError {

	err := csvc.cartRepository.Delete(customer_id)

	if err != nil {
		return err
	}

	return err
}

func (csvc CartServiceImpl) DeleteCartItem(customer_id string, product_id string) *app_error.AppError {

	// Get current cart object
	curr_cart, err := csvc.cartRepository.Read(customer_id)

	if err != nil {
		return err
	}

	// Find the product in the cart
	// Delete the product
	for i, p := range curr_cart.Products {
		if p.ProductId == product_id {
			curr_cart.Products = append(curr_cart.Products[:i], curr_cart.Products[i+1:]...)
		}
	}

	err2 := csvc.cartRepository.UpdateExisting(curr_cart)

	if err2 != nil {
		return err2
	}

	return nil
}

func (csvc CartServiceImpl) UserIDFromAuthToken(authToken string) (string, *app_error.AppError) {

	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(db.EnvJWTSecretKey()), nil
	})
	if err != nil {
		return "", app_error.NewAuthenticationError("unexpected signing method")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims["user_id"].(string), nil
	}
	return "", app_error.NewAuthenticationError("invalid token")
}

package config

import (
	"gatewayService/errs"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type EnvConfig struct {
	Port                      string
	AdminRouteKeywords        []string
	AuthGrpcHost              string
	AuthPathPrefix            string
	AuthGinPort               string
	AuthGrpcPort              string
	ProductsAdminPathPrefix   string
	ProductsAdminGinPort      string
	ProductsPathPrefix        string
	ProductsGinPort           string
	CustomerAdminPathPrefix   string
	CustomerAdminGinPort      string
	AccountPathPrefix         string
	AccountGinPort            string
	TransactionPathPrefix     string
	TransactionGinPort        string
	PaymentPathPrefix         string
	PaymentGinPort            string
	CartPathPrefix            string
	CartGinPort               string
	OrderPathPrefix           string
	OrderGinPort              string
	CategoriesPathPrefix      string
	CategoriesGinPort         string
	RewardsPathPrefix         string
	RewardsGinPort            string
	CheckoutPathPrefix        string
	CheckoutGinPort           string
	ShippingAddressPathPrefix string
	ShippingAddressGinPort    string
}

var EnvVars = &EnvConfig{}

func LoadEnvConfig() *errs.AppError {

	if err := godotenv.Load("./.env", "./local.env"); err != nil {
		return errs.NewUnexpectedError(err.Error())
	}

	EnvVars.Port = os.Getenv("PORT")
	EnvVars.AdminRouteKeywords = strings.Split(os.Getenv("ADMIN_ROUTE_KEYWORDS"), ",")

	EnvVars.AuthGrpcHost = os.Getenv("AUTH_GRPC_HOST")
	EnvVars.AuthPathPrefix = os.Getenv("AUTH_PATH_PREFIX")
	EnvVars.AuthGinPort = os.Getenv("AUTH_GIN_PORT")
	EnvVars.AuthGrpcPort = os.Getenv("AUTH_GRPC_PORT")
	EnvVars.ProductsAdminPathPrefix = os.Getenv("PRODUCTS_ADMIN_PATH_PREFIX")
	EnvVars.ProductsAdminGinPort = os.Getenv("PRODUCTS_ADMIN_GIN_PORT")
	EnvVars.ProductsPathPrefix = os.Getenv("PRODUCTS_PATH_PREFIX")
	EnvVars.ProductsGinPort = os.Getenv("PRODUCTS_GIN_PORT")
	EnvVars.CustomerAdminPathPrefix = os.Getenv("CUSTOMER_ADMIN_PATH_PREFIX")
	EnvVars.CustomerAdminGinPort = os.Getenv("CUSTOMER_ADMIN_GIN_PORT")
	EnvVars.AccountPathPrefix = os.Getenv("ACCOUNT_PATH_PREFIX")
	EnvVars.AccountGinPort = os.Getenv("ACCOUNT_GIN_PORT")
	EnvVars.TransactionPathPrefix = os.Getenv("TRANSACTION_PATH_PREFIX")
	EnvVars.TransactionGinPort = os.Getenv("TRANSACTION_GIN_PORT")
	EnvVars.PaymentPathPrefix = os.Getenv("PAYMENT_PATH_PREFIX")
	EnvVars.PaymentGinPort = os.Getenv("PAYMENT_GIN_PORT")
	EnvVars.CartPathPrefix = os.Getenv("CART_PATH_PREFIX")
	EnvVars.CartGinPort = os.Getenv("CART_GIN_PORT")
	EnvVars.OrderPathPrefix = os.Getenv("ORDER_PATH_PREFIX")
	EnvVars.OrderGinPort = os.Getenv("ORDER_GIN_PORT")
	EnvVars.CategoriesPathPrefix = os.Getenv("CATEGORIES_PATH_PREFIX")
	EnvVars.CategoriesGinPort = os.Getenv("CATEGORIES_GIN_PORT")
	EnvVars.RewardsPathPrefix = os.Getenv("REWARDS_PATH_PREFIX")
	EnvVars.RewardsGinPort = os.Getenv("REWARDS_GIN_PORT")
	EnvVars.CheckoutPathPrefix = os.Getenv("CHECKOUT_PATH_PREFIX")
	EnvVars.CheckoutGinPort = os.Getenv("CHECKOUT_GIN_PORT")
	EnvVars.ShippingAddressPathPrefix = os.Getenv("SHIPPING_ADDRESS_PATH_PREFIX")
	EnvVars.ShippingAddressGinPort = os.Getenv("SHIPPING_ADDRESS_GIN_PORT")

	return nil
}

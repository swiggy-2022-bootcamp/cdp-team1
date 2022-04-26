package routes

import (
	"gatewayService/app/handlers"
	"gatewayService/config"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterGatewayRoutes(gatewayRouter *mux.Router) {

	gatewayRouter.
		PathPrefix(config.EnvVars.AuthPathPrefix).
		Methods(http.MethodGet, http.MethodPost).
		Handler(http.HandlerFunc(handlers.NewGatewayHandler(config.EnvVars.AuthGinPort).ForwardRequest))
	gatewayRouter.
		PathPrefix(config.EnvVars.ProductsAdminPathPrefix).
		Methods(http.MethodGet, http.MethodPost).
		Handler(http.HandlerFunc(handlers.NewGatewayHandler(config.EnvVars.ProductsAdminGinPort).ForwardRequest))
	gatewayRouter.
		PathPrefix(config.EnvVars.ProductsPathPrefix).
		Methods(http.MethodGet, http.MethodPost).
		Handler(http.HandlerFunc(handlers.NewGatewayHandler(config.EnvVars.ProductsGinPort).ForwardRequest))
	gatewayRouter.
		PathPrefix(config.EnvVars.CustomerAdminPathPrefix).
		Methods(http.MethodGet, http.MethodPost).
		Handler(http.HandlerFunc(handlers.NewGatewayHandler(config.EnvVars.CustomerAdminGinPort).ForwardRequest))
	gatewayRouter.
		PathPrefix(config.EnvVars.AccountPathPrefix).
		Methods(http.MethodGet, http.MethodPost).
		Handler(http.HandlerFunc(handlers.NewGatewayHandler(config.EnvVars.AccountGinPort).ForwardRequest))
	gatewayRouter.
		PathPrefix(config.EnvVars.TransactionPathPrefix).
		Methods(http.MethodGet, http.MethodPost).
		Handler(http.HandlerFunc(handlers.NewGatewayHandler(config.EnvVars.TransactionGinPort).ForwardRequest))
	gatewayRouter.
		PathPrefix(config.EnvVars.PaymentPathPrefix).
		Methods(http.MethodGet, http.MethodPost).
		Handler(http.HandlerFunc(handlers.NewGatewayHandler(config.EnvVars.PaymentGinPort).ForwardRequest))
	gatewayRouter.
		PathPrefix(config.EnvVars.CartPathPrefix).
		Methods(http.MethodGet, http.MethodPost).
		Handler(http.HandlerFunc(handlers.NewGatewayHandler(config.EnvVars.CartGinPort).ForwardRequest))
	gatewayRouter.
		PathPrefix(config.EnvVars.OrderPathPrefix).
		Methods(http.MethodGet, http.MethodPost).
		Handler(http.HandlerFunc(handlers.NewGatewayHandler(config.EnvVars.OrderGinPort).ForwardRequest))
	gatewayRouter.
		PathPrefix(config.EnvVars.CategoriesPathPrefix).
		Methods(http.MethodGet, http.MethodPost).
		Handler(http.HandlerFunc(handlers.NewGatewayHandler(config.EnvVars.CategoriesGinPort).ForwardRequest))
	gatewayRouter.
		PathPrefix(config.EnvVars.RewardsPathPrefix).
		Methods(http.MethodGet, http.MethodPost).
		Handler(http.HandlerFunc(handlers.NewGatewayHandler(config.EnvVars.RewardsGinPort).ForwardRequest))
	gatewayRouter.
		PathPrefix(config.EnvVars.CheckoutPathPrefix).
		Methods(http.MethodGet, http.MethodPost).
		Handler(http.HandlerFunc(handlers.NewGatewayHandler(config.EnvVars.CheckoutGinPort).ForwardRequest))
	gatewayRouter.
		PathPrefix(config.EnvVars.ShippingAddressPathPrefix).
		Methods(http.MethodGet, http.MethodPost).
		Handler(http.HandlerFunc(handlers.NewGatewayHandler(config.EnvVars.ShippingAddressGinPort).ForwardRequest))
}

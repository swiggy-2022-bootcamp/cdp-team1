package app

import (
	"fmt"
	"gatewayService/app/middlewares"
	"gatewayService/app/routes"
	"gatewayService/config"
	"gatewayService/log"
	"github.com/gorilla/mux"
	"net/http"
)

func Start() {

	if fErr := config.LoadEnvConfig(); fErr != nil {
		log.Error("error in loading config :", fErr.Message)
	}
	port := config.EnvVars.Port

	middleware := middlewares.NewAuthMiddleware()

	router := mux.NewRouter()
	router.Use(middleware.VerifyRequest())

	routes.RegisterGatewayRoutes(router)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Error(err.Error())
	}

	log.Info("Starting gateway server on %s:%s ...\n\n", "127.0.0.1", port)
}

package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"gatewayService/config"
	"gatewayService/errs"
	"gatewayService/log"
	"gatewayService/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strings"
	"time"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (am *AuthMiddleware) VerifyRequest() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			reqMethod := r.Method
			reqUri := r.RequestURI
			fmt.Println(reqMethod, reqUri)
			if reqMethod == http.MethodPost && (strings.Contains(reqUri, "login") || strings.Contains(reqUri, "register")) {
				next.ServeHTTP(w, r)
			} else {

				role := "user"
				for _, key := range config.EnvVars.AdminRouteKeywords {
					if strings.Contains(reqUri, key) {
						role = "admin"
					}
				}
				if strings.HasPrefix(reqUri, config.EnvVars.AuthPathPrefix) {
					role = ""
				}

				token := r.Header.Get("Authorization")
				if token != "" {

					// Set up a connection to the server.
					serverAddress := fmt.Sprintf("%s:%s", config.EnvVars.AuthGrpcHost, config.EnvVars.AuthGrpcPort)
					conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
					if err != nil {
						log.Error("could not connect to Auth gRPC Server: ", err)
						return
					}
					//close
					defer func(conn *grpc.ClientConn) {
						err := conn.Close()
						if err != nil {
							log.Error("connection to Auth gRPC Server closed with error:", err.Error())
						}
					}(conn)
					c := protos.NewAuthClient(conn)

					// Disconnect gRPC call upon
					ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
					defer cancel()
					//var ctx context.Context

					authReq := &protos.TokenVerificationRequest{
						Token: token,
						Role:  role,
					}

					// Send gRPC request to server
					if authResp, err := c.VerifyToken(ctx, authReq); err != nil || authResp.GetIsVerified() == false {
						if int(authResp.GetStatusCode()) != 0 {
							w.WriteHeader(int(authResp.GetStatusCode()))
						} else {
							w.WriteHeader(http.StatusInternalServerError)
						}

						respBody := errs.AppError{
							Message: authResp.GetStatusMessage(),
						}
						respJson, _ := json.Marshal(respBody.AsMessage())
						w.Write(respJson)
						return
					} else {
						next.ServeHTTP(w, r)
					}
				} else {
					err := errs.NewValidationError("Authorization token missing from header")
					w.WriteHeader(err.Code)
					respJson, _ := json.Marshal(err.AsMessage())
					w.Write(respJson)
				}
			}
		})
	}
}

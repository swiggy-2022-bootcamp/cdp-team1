package domain

import (
	"authService/protos"
	"context"
)

type AuthProtoService struct {
	AuthSvc AuthService
	protos.UnimplementedAuthServer
}

func NewAuthProtoService(authSvc AuthService) *AuthProtoService {
	return &AuthProtoService{
		AuthSvc: authSvc,
	}
}

func (aps AuthProtoService) VerifyToken(ctx context.Context, request *protos.TokenVerificationRequest) (*protos.TokenVerificationResponse, error) {

	response := &protos.TokenVerificationResponse{}
	isValid, err := aps.AuthSvc.VerifyAuthToken(request.GetToken(), "", request.GetRole())

	response.IsVerified = isValid
	if err != nil {
		response.StatusCode = int32(err.Code)
		response.StatusMessage = err.Message
	}
	return response, nil
}

package authorization_handler

import auth "github.com/Fitray/auth_service/api/gen/go"

type AuthorizationHandler struct {
	auth.UnimplementedAuthorizationServiceServer
	service AuthorizationService
}

type AuthorizationService interface{}

func NewAuthorizationHandler(
	service AuthorizationService,
) *AuthorizationHandler {
	return &AuthorizationHandler{
		service: service,
	}
}

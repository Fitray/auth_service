package authorization_service

type AuthorizationService struct {
	authorizationPostgres AuthorizationPostgres
	authorizationRedis    AuthorizationRedis
}

type AuthorizationPostgres interface{}
type AuthorizationRedis interface{}

func NewAuthorizationService(
	authorizationPostgres AuthorizationPostgres,
	authorizationRedis AuthorizationRedis,
) *AuthorizationService {
	return &AuthorizationService{
		authorizationPostgres: authorizationPostgres,
		authorizationRedis:    authorizationRedis,
	}
}

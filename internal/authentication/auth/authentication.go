package auth

import (
	"context"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
)

type Authentication interface {
	SignUp(ctx context.Context, name, username, password, role string) (*models.AuthTokens, error)
	Login(ctx context.Context, username, password string) (*models.AuthTokens, error)
	Refresh(ctx context.Context, refreshToken string) (*models.AuthTokens, error)
	ValidateAccessToken(ctx context.Context, token string) (*models.AuthUser, error)
	Logout(ctx context.Context, id string) error
}

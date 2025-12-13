package jwt

import (
	"context"
	"fmt"
	"time"

	"github.com/SunilKividor/PillNet-Backend/internal/authentication/auth"
	"github.com/SunilKividor/PillNet-Backend/internal/db/repository"
	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/golang-jwt/jwt"
)

const issuer = "PillNet"

type JWTAuthClient struct {
	authRepo   *repository.AuthRepository
	privateKey []byte
}

func NewJWTAuthenticationClient(authRepo *repository.AuthRepository, jwtSecret string) *JWTAuthClient {

	return &JWTAuthClient{
		authRepo:   authRepo,
		privateKey: []byte(jwtSecret),
	}
}

func (a *JWTAuthClient) SignUp(ctx context.Context, name, username, password, role string) (*models.AuthTokens, error) {

	passwordHash, err := auth.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error hashing password")
	}

	userId, err := a.authRepo.RegisterNewUser(ctx, name, username, passwordHash, role)
	if err != nil {
		return nil, fmt.Errorf("no user found")
	}

	refreshToken, err := a.generateRefreshToken(userId)
	if err != nil {
		return nil, err
	}
	accessToken, err := a.generateAccessToken(userId)
	if err != nil {
		return nil, err
	}

	expTime := time.Now().Add(30 * 24 * time.Hour)
	refreshTokenExp := time.Until(expTime)
	if a.authRepo.SetRefreshTokenS(ctx, userId, refreshToken, refreshTokenExp) != nil {
		return nil, err
	}

	authTokens := &models.AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return authTokens, nil
}

func (a *JWTAuthClient) Login(ctx context.Context, username, password string) (*models.AuthTokens, error) {

	user, err := a.authRepo.GetUserByEmail(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("no user found")
	}

	if !auth.ComparePassword(user.PasswordHash, password) {
		return nil, fmt.Errorf("invalid credentials")
	}

	refreshToken, err := a.generateRefreshToken(user.Id)
	if err != nil {
		return nil, err
	}
	accessToken, err := a.generateAccessToken(user.Id)
	if err != nil {
		return nil, err
	}

	expTime := time.Now().Add(30 * 24 * time.Hour)
	refreshTokenExp := time.Until(expTime)
	if a.authRepo.SetRefreshTokenS(ctx, user.Id, refreshToken, refreshTokenExp) != nil {
		return nil, err
	}

	authTokens := &models.AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return authTokens, nil
}

func (a *JWTAuthClient) Refresh(ctx context.Context, refreshToken string) (*models.AuthTokens, error) {

	parsedToken, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Method.Alg())
		}
		return a.privateKey, nil
	})

	if err != nil || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid refresh token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	userId, ok := claims["sub"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid sub claim")
	}

	exp := int64(claims["exp"].(float64))
	if time.Now().Unix() > exp {
		return nil, fmt.Errorf("refresh token expired")
	}

	storedRefreshToken, err := a.authRepo.GetRefreshToken(ctx, userId)
	if err != nil || refreshToken != storedRefreshToken {
		return nil, fmt.Errorf("invalid refresh token")
	}

	accessToken, err := a.generateAccessToken(userId)
	if err != nil {
		return nil, err
	}

	return &models.AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: storedRefreshToken,
	}, nil
}

func (a *JWTAuthClient) ValidateAccessToken(ctx context.Context, token string) (*models.AuthUser, error) {
	return nil, nil
}

func (a *JWTAuthClient) Logout(ctx context.Context, id string) error {

	if id == "" {
		return fmt.Errorf("empty id")
	}

	err := a.authRepo.DeleteRefreshToken(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

//----------------------------------------utils----------------------------------------

func (a *JWTAuthClient) generateAccessToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"iss": issuer,
		"sub": userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.privateKey)
}

func (a *JWTAuthClient) generateRefreshToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"iss": issuer,
		"sub": userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(30 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.privateKey)
}

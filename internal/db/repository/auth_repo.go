package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type AuthRepository struct {
	db *pgxpool.Pool
	rd *redis.Client
}

func NewAuthRepository(db *pgxpool.Pool, rd *redis.Client) *AuthRepository {
	return &AuthRepository{
		db: db,
		rd: rd,
	}
}

func refreshKey(id string) string {
	return fmt.Sprintf("refresh:%s", id)
}

func (a *AuthRepository) GetUserByEmail(ctx context.Context, username string) (*models.AuthUser, error) {

	db := a.db

	smt := `SELECT id,username,password_hash,role FROM users WHERE username = $1`

	var authUser models.AuthUser
	row := db.QueryRow(ctx, smt, username)
	err := row.Scan(&authUser.Id, &authUser.Username, &authUser.PasswordHash, &authUser.Role)
	if err != nil {
		return nil, err
	}

	return &authUser, nil
}

func (a *AuthRepository) RegisterNewUser(ctx context.Context, name, username, password, role string) (string, error) {

	db := a.db

	smt := `INSERT INTO users(name,username,password_hash,role) VALUES($1,$2,$3,$4) RETURNING id`

	var id uuid.UUID
	err := db.QueryRow(ctx, smt, name, username, password, role).Scan(&id)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (a *AuthRepository) SetRefreshTokenS(ctx context.Context, id, refreshToken string, exp time.Duration) error {
	redisClient := a.rd

	key := refreshKey(id)

	err := redisClient.Set(ctx, key, refreshToken, exp).Err()
	if err != nil {
		return fmt.Errorf("failed to store refresh token: %w", err)
	}

	return nil
}

func (a *AuthRepository) GetRefreshToken(ctx context.Context, id string) (string, error) {
	redisClient := a.rd

	key := refreshKey(id)

	refreshToken, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	if refreshToken == "" {
		return "", fmt.Errorf("empty key")
	}

	return refreshToken, nil
}

func (a *AuthRepository) DeleteRefreshToken(ctx context.Context, id string) error {
	redisClient := a.rd

	key := refreshKey(id)

	res, err := redisClient.Del(ctx, key).Result()

	if res == 0 {
		return fmt.Errorf("refresh token does not exist or already deleted")
	}

	if err != nil {
		return fmt.Errorf("failed to delete refresh token: %w", err)
	}

	return nil
}

package di

import (
	"github.com/SunilKividor/PillNet-Backend/internal/api"
	"github.com/SunilKividor/PillNet-Backend/internal/authentication/http/middleware"
	"github.com/SunilKividor/PillNet-Backend/internal/authentication/jwt"
	"github.com/SunilKividor/PillNet-Backend/internal/config"
	"github.com/SunilKividor/PillNet-Backend/internal/db/pg"
	redisdb "github.com/SunilKividor/PillNet-Backend/internal/db/redis"
	"github.com/SunilKividor/PillNet-Backend/internal/db/repository"
	"github.com/SunilKividor/PillNet-Backend/internal/handler"
)

func InitializeApp() (*api.Server, error) {
	// ctx := context.Background()

	cfg := config.Load()

	server := api.NewServer(cfg)

	pgConn := pg.NewConnection(cfg.PostgresConfig.ConnectionString)
	pool, err := pgConn.Connect()
	if err != nil {
		return nil, err
	}

	redisConn := redisdb.NewConnection(cfg.RedisConfig.ConnectionString)
	redisClient, err := redisConn.Connect()
	if err != nil {
		return nil, err
	}

	jwtRepo := repository.NewAuthRepository(pool, redisClient)
	jwtAuth := jwt.NewJWTAuthenticationClient(jwtRepo, cfg.JWTConfig.Secret)

	handlers := &handler.Handlers{
		Authentication: handler.NewAuthenticationHandler(jwtAuth),
	}
	middleware := middleware.JWTMiddleware()

	api.RegisterRoutes(server.Engine, cfg, handlers, middleware)

	return server, nil
}

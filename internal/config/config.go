package config

import "os"

type Config struct {
	ServerConfig   *ServerConfig
	PostgresConfig *PostgresConfig
	JWTConfig      *JWTConfig
	RedisConfig    *RedisConfig
}

type ServerConfig struct {
	Port string
}

type PostgresConfig struct {
	ConnectionString string
}

type JWTConfig struct {
	Secret string
}

type RedisConfig struct {
	ConnectionString string
}

func Load() *Config {
	return &Config{
		ServerConfig: &ServerConfig{
			Port: os.Getenv("PORT"),
		},
		PostgresConfig: &PostgresConfig{
			ConnectionString: os.Getenv("Postgres_URI"),
		},
		RedisConfig: &RedisConfig{
			ConnectionString: os.Getenv("Redis_URI"),
		},
		JWTConfig: &JWTConfig{
			Secret: os.Getenv("JWTAPISECRET"),
		},
	}
}

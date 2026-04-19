package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Addr string `env:"HTTP_ADDR" envDefault:":8080"`
}

type DatabaseConfig struct {
	Name     string `env:"DB_NAME,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Port     int    `env:"DB_PORT,required"`
	Host     string `env:"DB_HOST,required"`
	SSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
}

type JWTConfig struct {
	Secret     string `env:"JWT_SECRET,required"`
	AccessTTL  int    `env:"JWT_ACCESS_TTL" envDefault:"10"`
	RefreshTTL int    `env:"JWT_REFRESH_TTL" envDefault:"15"`
}

func Load() (config *Config, err error) {
	err = godotenv.Load()
	if err != nil {
		return nil, domain.ErrEnvNotFound
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validate config: %w", err)
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	if len(c.JWT.Secret) < 32 {
		return domain.ErrShortJwtSecret
	}

	if c.JWT.AccessTTL > 15 {
		return domain.ErrJwtAccessLong
	}

	if c.JWT.RefreshTTL > 30 {
		return domain.ErrJwtRefreshLong
	}
	return nil
}

func (d *DatabaseConfig) GenerateDsn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host,
		d.Port,
		d.User,
		d.Password,
		d.Name,
		d.SSLMode,
	)
}

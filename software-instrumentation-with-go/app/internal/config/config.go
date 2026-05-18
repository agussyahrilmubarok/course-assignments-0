package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		Name string `mapstructure:"name"`
	} `mapstructure:"server"`

	Postgres struct {
		Host            string `mapstructure:"host"`
		Port            int    `mapstructure:"port"`
		User            string `mapstructure:"user"`
		Password        string `mapstructure:"password"`
		Name            string `mapstructure:"name"`
		SSLMode         string `mapstructure:"sslmode"`
		Timezone        string `mapstructure:"timezone"`
		MaxIdleConns    int    `mapstructure:"max_idle_conns"`
		MaxOpenConns    int    `mapstructure:"max_open_conns"`
		ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
	} `mapstructure:"postgres"`
}

func Load(filepath string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(filepath)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := validate(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validate(cfg *Config) error {
	// Server
	if cfg.Server.Host == "" {
		return fmt.Errorf("server.host is required")
	}
	if cfg.Server.Port <= 0 {
		return fmt.Errorf("server.port must be greater than 0")
	}
	if cfg.Server.Name == "" {
		return fmt.Errorf("server.name is required")
	}

	// Postgres
	if cfg.Postgres.Host == "" {
		return fmt.Errorf("postgres.host is required")
	}
	if cfg.Postgres.Port <= 0 {
		return fmt.Errorf("postgres.port must be greater than 0")
	}
	if cfg.Postgres.User == "" {
		return fmt.Errorf("postgres.user is required")
	}
	if cfg.Postgres.Password == "" {
		return fmt.Errorf("postgres.password is required")
	}
	if cfg.Postgres.Name == "" {
		return fmt.Errorf("postgres.name is required")
	}
	if cfg.Postgres.SSLMode == "" {
		return fmt.Errorf("postgres.sslmode is required")
	}
	if cfg.Postgres.Timezone == "" {
		return fmt.Errorf("postgres.timezone is required")
	}
	if cfg.Postgres.MaxIdleConns < 0 {
		return fmt.Errorf("postgres.max_idle_conns must be >= 0")
	}
	if cfg.Postgres.MaxOpenConns < 0 {
		return fmt.Errorf("postgres.max_open_conns must be >= 0")
	}
	if cfg.Postgres.ConnMaxLifetime < 0 {
		return fmt.Errorf("postgres.conn_max_lifetime must be >= 0")
	}

	return nil
}

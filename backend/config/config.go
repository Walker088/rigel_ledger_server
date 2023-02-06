package config

import (
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
)

func init() {
	err := godotenv.Load(".env.development")
	if err != nil {
		log.Fatal("Error loading .env.development file")
	}
}

type AppConfig struct {
	AppHost         string   `env:"APP_HOST" envDefault:"localhost"`
	AppPort         string   `env:"APP_PORT" envDefault:"8000"`
	AppAllowOrigins []string `env:"APP_ALLOW_ORIGINS"`
	JwtSecret       string   `env:"JWT_SECRET"`

	GithubOAuthConfig
}

type PgConfig struct {
	DbHost   string `env:"DB_HOST" envDefault:"localhost"`
	DbPort   int    `env:"DB_PORT" envDefault:"5432"`
	DbSchema string `env:"DB_SCHEMA" envDefault:"public"`
	DbName   string `env:"DB_NAME" envDefault:"rigel_ledger"`
	DbUser   string `env:"DB_USER" envDefault:"postgres"`
	DbPass   string `env:"DB_PW"`

	MinConns              int32         `env:"DB_MIN_CONN" envDefault:"10"`
	MaxConns              int32         `env:"DB_MAX_CONN" envDefault:"100"`
	MaxConnIdleTime       time.Duration `env:"DB_MAX_CONN_IDLE" envDefault:"10m"`
	MaxConnLifetime       time.Duration `env:"DB_MAX_CONN_LIFETIME" envDefault:"30m"`
	MaxConnLifetimeJitter time.Duration `env:"DB_MAX_CONN_LIFETIME_JITTER" envDefault:"1m"`
	HealthCheckPeriod     time.Duration `env:"DB_HEALTH_CHECK_PERIOD" envDefault:"1m"`
}

type LoggerConfig struct {
	Level         string `env:"LOG_LEVEL" envDefault:"info"`
	EncoderConfig zapcore.EncoderConfig
}

type GithubOAuthConfig struct {
	OauthGithubClientId     string `env:"OATH_GITHUB_CLIENT_ID"`
	OauthGithubClientSecret string `env:"OATH_GITHUB_CLIENT_SECRET"`
}

func (pg *PgConfig) ToConnString() string {
	//urlExample := postgres://username:password@localhost:5432/database_name
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		pg.DbUser,
		pg.DbPass,
		pg.DbHost,
		pg.DbPort,
		pg.DbName,
	)
}

func GetAppConfig() *AppConfig {
	cfg := AppConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}
	return &cfg
}

func GetPgConfig() *PgConfig {
	cfg := PgConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}
	return &cfg
}

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(
		fmt.Sprintf(
			"%d-%02d-%02d %02d:%02d:%02d",
			t.Year(),
			t.Month(),
			t.Day(),
			t.Hour(),
			t.Minute(),
			t.Second(),
		),
	)
}

func GetLoggerConfig() *LoggerConfig {
	encoder := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     formatEncodeTime,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	logCfg := LoggerConfig{EncoderConfig: encoder}
	if err := env.Parse(&logCfg); err != nil {
		log.Fatalf("%+v\n", err)
	}
	return &logCfg
}

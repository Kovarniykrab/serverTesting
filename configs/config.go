package configs

import "log/slog"

type Config struct {
	PSQL
	JWT
	Web
}

type JWT struct {
	Issuer      string `long:"jwt-issuer" env:"SERVERTESTING_JWT_ISSUER"`
	SecretKey   string `long:"jwt-secret-key" env:"SERVERTESTING_JWT_SECRET_KEY"`
	KeyLength   int    `long:"jwt-key-length" env:"SERVERTESTING_JWT_KEY_LENGTH"`
	HourExpired int    `long:"jwt-hour-expired" env:"SERVERTESTING_JWT_HOUR_EXPIRED"`
}
type Web struct {
	Host         string     `long:"web-host" env:"SERVERTESTING_WEB_HOST"`
	Port         int        `long:"web-port" env:"SERVERTESTING_WEB_PORT"`
	ReadTimeout  int        `long:"web-read-timeout" env:"SERVERTESTING_WEB_READ_TIMEOUT"`
	WriteTimeout int        `long:"web-write-timeout" env:"SERVERTESTING_WEB_WRITE_TIMEOUT"`
	IdleTimeout  int        `long:"web-idle-timeout" env:"SERVERTESTING_WEB_IDLE_TIMEOUT"`
	SSLSertPath  string     `long:"web-ssl-sert-path" env:"SERVERTESTING_WEB_SSL_SERT_PATH"`
	SSLKeyPath   string     `long:"web-ssl-key-path" env:"SERVERTESTING_WEB_SSL_KEY_PATH"`
	LogLevel     slog.Level `long:"web-log-level" env:"SERVERTESTING_WEB_LOG_LEVEL"`
}
type PSQL struct {
	Host     string `long:"psql-host" env:"SERVERTESTING_PSQL_HOST"`
	Port     int    `long:"psql-port" env:"SERVERTESTING_PSQL_PORT"`
	DSN      string `long:"psql-dsn" env:"SERVERTESTING_PSQL_DSN" `
	UserName string `long:"psql-user-name" env:"SERVERTESTING_PSQL_USER_NAME"`
	DBName   string `long:"psql-db-name" env:"SERVERTESTING_PSQL_DB_NAME"`
	Password string `long:"psql-password" env:"SERVERTESTING_PSQL_PASSWORD"`
	SslMode  string `long:"psql-ssl-mode" env:"SERVERTESTING_PSQL_SSL_MODE"`
}

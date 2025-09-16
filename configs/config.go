package configs

import "log/slog"

type Config struct {
	PSQL
	JWT
	Web
}

type JWT struct {
	SecretKey string `long:"jwt-secret-key" env:"SERVER_JWT_SECRET_KEY"`
}
type Web struct {
	Host         string     `long:"web-host" env:"SERVER_WEB_HOST"`
	Port         int        `long:"web-port" env:"SERVER_WEB_PORT"`
	ReadTimeout  int        `long:"web-read-timeout" env:"SERVER_WEB_READ_TIMEOUT"`
	WriteTimeout int        `long:"web-write-timeout" env:"SERVER_WEB_WRITE_TIMEOUT"`
	IdleTimeout  int        `long:"web-idle-timeout" env:"SERVER_WEB_IDLE_TIMEOUT"`
	Cors         string     `long:"web-cors" env:"SERVER_WEB_CORS"`
	SSLSertPath  string     `long:"web-ssl-sert-path" env:"SERVER_WEB_SSL_SERT_PATH"`
	SSLKeyPath   string     `long:"web-ssl-web" env:"SERVER_WEB_SSL_KEY_PATH"`
	LogLevel     slog.Level `long:"web-log-level" env:"SERVER_WEB_LOG_LEVEL"`
}
type PSQL struct {
	DSN string `long:"psql-dsn" env:"SERVER_PSQL_DSN" `
}

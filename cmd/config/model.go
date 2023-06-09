package config

import (
	"ssr/internal/infra/markepplacerepo"
)

type ServerConfig struct {
	Port int `json:"port"`
}

type Config struct {
	Server           ServerConfig                `json:"server"`
	MarketplaceMySql markepplacerepo.MysqlConfig `json:"marketplace_db"`
}

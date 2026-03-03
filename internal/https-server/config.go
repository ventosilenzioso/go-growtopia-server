package httpsserver

import (
	"os"
	"github.com/BurntSushi/toml"
)

type Config struct {
	ENet  ENetConfig  `toml:"enet"`
	HTTPS HTTPSConfig `toml:"https"`
}

type ENetConfig struct {
	Port     int `toml:"port"`
	MaxPeers int `toml:"max_peers"`
}

type HTTPSConfig struct {
	Port     int    `toml:"port"`
	HTTPPort int    `toml:"http_port"`
	CertFile string `toml:"cert_file"`
	KeyFile  string `toml:"key_file"`
}

func LoadConfig() (*Config, error) {
	var config Config
	data, err := os.ReadFile("ServerConfiguration.toml")
	if err != nil {
		return nil, err
	}
	
	if err := toml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	
	return &config, nil
}

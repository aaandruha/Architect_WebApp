package config

import "github.com/kelseyhightower/envconfig"

const appID = "orderservice"

type config struct {
	ServeRESTAddress string `envconfig:"server_rest_address" default:":8000"`
	DBName           string `envconfig:"server_db_address" default:"orderservice"`
	DBHost           string `envconfig:"server_db_host" default:":3306"`
	DBUser           string `envconfig:"server_db_user" default:"root"`
	DBPassword       string `envconfig:"server_db_password" default:""`
}

func ParseEnv() (*config, error) {
	c := new(config)
	if err := envconfig.Process(appID, c); err != nil {
		return nil, err
	}
	return c, nil
}

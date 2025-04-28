package conf

import (
	"github.com/llmons/who-is-the-undercover/internal/base/data"
	"github.com/llmons/who-is-the-undercover/internal/base/server"
	"github.com/spf13/viper"
)

type AllConfig struct {
	Debug  bool    `json:"debug" yaml:"debug"`
	Server *Server `json:"server" yaml:"server"`
	Data   *Data   `json:"data" yaml:"data"`
}

type Server struct {
	HTTP *server.HTTP `json:"http" yaml:"http"`
}

type Data struct {
	Database *data.Database `json:"database" yaml:"database"`
}

func ReadConfig(configFilePath string) (c *AllConfig, err error) {
	v := viper.New()
	v.SetConfigFile(configFilePath)

	if err = v.ReadInConfig(); err != nil {
		return nil, err
	}

	c = &AllConfig{}
	if err = v.Unmarshal(c); err != nil {
		return nil, err
	}

	return c, nil
}

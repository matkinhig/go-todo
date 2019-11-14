package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Schema struct {
	OracleDB struct {
		Uri      string `mapstructure:"Uri"`
		Host     string `mapstructure:"Host"`
		Username string `mapstructure:"User"`
		Password string `mapstructure:"Password"`
	} `mapstructure:"OracleDB"`

	Encryption struct {
		JWTSecret string
	}
}

var Config Schema

// [username/[password]@]host[:port][/service_name][?param1=value1&...&paramN=valueN]

func init() {
	viper.SetConfigName("config") // name of config file (without extension)

	viper.AddConfigPath("./config") // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v\n", err))
	}
}

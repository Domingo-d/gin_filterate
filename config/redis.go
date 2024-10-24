package config

type (
	RedisConfig struct {
		Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
		Password string `mapstructure:"password" json:"password" yaml:"password"`
		DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	}
)

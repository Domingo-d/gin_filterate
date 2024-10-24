package config

type (
	ServerConfig struct {
		System        SystemConfig `mapstructure:"system" json:"system" yaml:"system"`
		RedisConf     RedisConfig  `mapstructure:"redisConf" json:"redisConf" yaml:"redisConf"`
		FilterateName string       `mapstructure:"filterateName" json:"filterateName" yaml:"filterateName"`
		ZapConf       ZapConf      `mapstructure:"zapConf" json:"zapConf" yaml:"zapConf"`
		JwtKey        string       `mapstructure:"jwtKey" json:"jwtKey" yaml:"jwtKey"`
		CertFile      string       `mapstructure:"certFile" json:"certFile" yaml:"certFile"`
		KeyFile       string       `mapstructure:"keyFile" json:"keyFile" yaml:"keyFile"`
	}
)

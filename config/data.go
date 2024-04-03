package config

type Config struct {
	Server ConfigServer `mapstructure:"server"`
	DB     ConfigDB     `mapstructure:"db"`
	Logger ConfigLogger `mapstructure:"logger"`
	JWT    ConfigJWT    `mapstructure:"jwt"`
}

type ConfigServer struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
}

type ConfigDB struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Member   string `mapstructure:"member"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type ConfigLogger struct {
	Level string `mapstructure:"level"`
}

type ConfigJWT struct {
	Secret     string `mapstructure:"secret"`
	Expiration int    `mapstructure:"expiration"`
}

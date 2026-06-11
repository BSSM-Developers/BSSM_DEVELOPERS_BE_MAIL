package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	SMTP   SMTPConfig
	Auth   AuthConfig
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type SMTPConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	From     string `mapstructure:"from"`
}

type AuthConfig struct {
	Secret string `mapstructure:"secret"`
}

func bindEnvs(v *viper.Viper) {
	keys := []string{
		"server.port",
		"smtp.host", "smtp.port", "smtp.username", "smtp.password", "smtp.from",
		"auth.secret",
	}
	for _, k := range keys {
		_ = v.BindEnv(k)
	}
}

func Load() *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	v.SetEnvPrefix("MAIL")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	bindEnvs(v)

	v.SetDefault("server.port", "8091")
	v.SetDefault("smtp.host", "smtp.gmail.com")
	v.SetDefault("smtp.port", 587)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic("config load error: " + err.Error())
		}
	}

	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		panic("config unmarshal error: " + err.Error())
	}
	return cfg
}

package config

import (
	"fmt"
	"log"
	"ys_go/config"

	"github.com/spf13/viper"
)

func InitConfig() *config.Config {
	configName := "config"

	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置失败 [%s.yaml]: %v", configName, err)
	}

	cfg := &config.Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("配置文件解析失败: %v", err)
	}
	fmt.Println("配置加载成功:", viper.ConfigFileUsed())
	return cfg
}

package main

import (
	"github.com/spf13/viper"
	"log"
)

type ConfigurationFactory struct {
	vConfig *viper.Viper
}

func (c *ConfigurationFactory) VConfig() *viper.Viper {
	return c.vConfig
}

func (c *ConfigurationFactory) SetVConfig(vConfig *viper.Viper) {
	c.vConfig = vConfig
}

func NewConfigurationFactory() *ConfigurationFactory {
	v := viper.New()
	v.SetConfigFile("./config.toml")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	return &ConfigurationFactory{
		vConfig: v,
	}
}

func (c *ConfigurationFactory) NewKafkaConfiguration() *KafkaConfiguration {
	return NewKafkaConfiguration(
		c.VConfig().GetString("kafka.key"),
		c.VConfig().GetString("kafka.topic"),
		c.VConfig().GetString("kafka.groupid"),
		c.VConfig().GetString("kafka.ip"),
		c.VConfig().GetString("kafka.port"),
	)
}

func (c *ConfigurationFactory) NewConfiguration() *Configuration {
	return NewConfiguration(c.VConfig().GetInt("execute.loop"), c.VConfig().GetString("execute.input_filename"))
}

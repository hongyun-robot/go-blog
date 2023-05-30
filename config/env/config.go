package env

import (
	"fmt"
	"gopkg.in/ini.v1"
)

func InitConfig() *Config {
	return getBaseConfig()
}

func getBaseConfig() *Config {

	// 基础信息
	baseInfo := BaseInfo{env: "dev"}

	fmt.Println(baseInfo.env, "baseConfig.env")
	// 读取配置文件路径
	path := getAppPath() + "\\env\\" + baseInfo.env + ".ini"

	// 无默认值
	// config := &Config{BaseConfig: BaseConfig{}}

	// 设置默认值
	config := &Config{BaseConfig: BaseConfig{
		ENV:           "pro",
		HTTP_PORT:     8888,
		READ_TIMEOUT:  60,
		WRITE_TIMEOUT: 60,
	}}

	// 映射，一切竟可以如此的简单。
	err := ini.MapTo(config, path)
	if err != nil {
		fmt.Println("ini文件映射出错啦，错误是:", err)
	}

	fmt.Println(config, config.BaseConfig, "baseConfig")

	return config
}

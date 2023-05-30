package env

type BaseInfo struct {
	env string
}

type Config struct {
	BaseConfig `ini:"BASE_CONFIG"`
	LogConfig  `ini:"LOG_CONFIG"`
}

type BaseConfig struct {
	ENV           string `ini:"ENV"`
	HTTP_PORT     int    `ini:"HTTP_PORT"`
	READ_TIMEOUT  int    `ini:"READ_TIMEOUT"`
	WRITE_TIMEOUT int    `ini:"WRITE_TIMEOUT"`
}

type LogConfig struct {
	LOG_SAVE_URL  string `ini:"LOG_SAVE_URL"`
	LOG_SAVE_NAME string `ini:"LOG_SAVE_NAME"`
}

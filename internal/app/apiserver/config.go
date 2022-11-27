package apiserver

//Конфиг. Данные для старта сервера. Порт и т.д.
type Config struct {
	BindAddr string `toml:"bind_add"`
	LogLevel string `toml:"log_level"`
}

//Функция которая возвращает новый конфиг
func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		LogLevel: "debug",
	}
}

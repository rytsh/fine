package config

var (
	App = Config{
		Log: Log{
			Level: "info",
		},
		Server: Server{
			Host: "0.0.0.0",
			Port: 8080,
		},
	}

	File = "./config.yaml"
)

type Config struct {
	Log     Log     `cfg:"log"`
	Server  Server  `cfg:"server"`
	Storage Storage `cfg:"storage"`
}

type Log struct {
	Level string `cfg:"level"`
}

type Server struct {
	Host     string `cfg:"host"`
	Port     int    `cfg:"port"`
	BasePath string `cfg:"basePath"`
}

type Storage struct {
	Local *Local `cfg:"local"`
}

type Local struct {
	Path string `cfg:"path"`
}

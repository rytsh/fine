package config

type AppInfo struct {
	Version     string
	BuildCommit string
	BuildDate   string
}

var (
	Info AppInfo
	Name string = "fine"
)

func SetInfo(info AppInfo) {
	Info = info
}

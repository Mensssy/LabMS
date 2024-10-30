package conf

var (
	ServerPort string

	TokenKey string
)

func InitWebConf() {
	ServerPort = ":80"
	TokenKey = ""
}

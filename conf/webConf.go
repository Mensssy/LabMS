package conf

var (
	ServerPort string

	TokenKey string
)

func InitWebConf() {
	ServerPort = ":8080"
	TokenKey = "CreamMushroomPasta"
}

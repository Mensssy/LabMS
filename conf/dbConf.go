package conf

var (
	SqlUserName string
	SqlPassword string
	SqlProtocol string
	SqlHost     string
	SqlPort     string
	SqlName     string
	SqlParam    string
)

func InitDBConf() {
	initSql()
	initMongo()
}

func initMongo() {
}

func initSql() {
	SqlUserName = "root"
	SqlPassword = "02040608"
	// SqlPassword = "BurgerSteak"
	SqlProtocol = "@tcp"
	SqlHost = "127.0.0.1"
	// SqlHost = "http://120.46.36.230"
	SqlPort = "3306"
	SqlName = "labms"
	SqlParam = "charset=utf8mb4&parseTime=True&loc=Local"
}

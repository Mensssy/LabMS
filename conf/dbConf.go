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
	SqlPassword = ""
	SqlProtocol = "@tcp"
	SqlHost = "127.0.0.1"
	SqlPort = "3306"
	SqlName = "LabMS"
	SqlParam = "charset=utf8mb4&parseTime=True"
}

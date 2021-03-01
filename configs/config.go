package configs

import (
	"testAPI/utils/environtment"
)

type dbConf struct {
	DbUser string
	DbPass string
	DbHost string
	DbPort string
	DbSchema string
	AllowNativePasswords bool
}
type Conf struct {
	Db dbConf
}

func NewAppConfig() *Conf  {
	return &Conf{dbConf{
		DbUser: environtment.Get("DbUser", "root"),
		DbPass: environtment.Get("DbPass", "pass"),
		DbHost: environtment.Get("DbHost", "localhost"),
		DbPort: environtment.Get("DbPort", "3306"),
		DbSchema: environtment.Get("DbSchema", "linkaja"),
		AllowNativePasswords: true,
	}}
}

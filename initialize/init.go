package initialize

import (
	"github.com/Dream-ming/myMusic/internal/config"
)

func InitAll() {
	
	config.LoadOSSConfig()
	config.LoadMySQLConfig()
	config.LoadJWTConfig()
	
	oss_init()
	mysql_init()
}
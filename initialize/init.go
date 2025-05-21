package initialize

import (
	"github.com/Dream-ming/myMusic/internal/config"
)

func InitAll() {
	
	config.LoadOSSConfig()
	
	oss_init()
	mysql_init()
}
package initialize

import (
	"github.com/Dream-ming/myMusic/internal/config"
)

func InitAll() {
	
	config.LoadOSSConfig()
	initoss()
}
package initialize

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func mysql_init() {
    // 格式: "用户名:密码@tcp(IP:端口)/数据库名称?charset=utf8&parseTime=True"
    dsn := "root:kedaya@tcp(127.0.0.1:3306)/myMusic?charset=utf8mb4&parseTime=True"
    
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(fmt.Errorf("open mysql failed: %w", err))
    }
    if err = db.Ping(); err != nil {
        panic(fmt.Errorf("ping mysql failed: %w", err))
    }
    DB = db
    fmt.Println("MySQL connected.")
}
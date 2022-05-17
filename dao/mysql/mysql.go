package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"web_app/settings"
)

//初始化数据库
var db *sqlx.DB

func Init(cfg *settings.MysqlConfig) (err error) {
	//读取配置文件
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("数据库连接失败，err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(cfg.MaxConnections)
	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	return
}

// 关闭链接

func Close() {
	_ = db.Close()
}

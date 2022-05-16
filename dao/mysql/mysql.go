package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

//初始化数据库
var db *sqlx.DB

func Init() (err error) {
	//读取配置文件
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"))

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("数据库连接失败，err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_connections"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_connections"))
	return
}

// 关闭链接

func Close() {
	_ = db.Close()
}

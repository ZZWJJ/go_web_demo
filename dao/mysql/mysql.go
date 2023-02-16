package mysql

import (
	"fmt"
	"web_app/settings"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(MysqlConfig *settings.MysqlConfig) (err error) {
	// DSN:Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		MysqlConfig.UserName,
		MysqlConfig.Password,
		MysqlConfig.Host,
		MysqlConfig.Port,
		MysqlConfig.DbName,
	)

	// zap.L().Sugar().Infof("max_open_conns: %d", viper.GetInt("mysql.max_open_conns"))
	// zap.L().Sugar().Infof("max_idle_conns: %d", viper.GetInt("mysql.max_idle_conns"))

	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sqlx.Connect("mysql", dsn)
	db.SetMaxOpenConns(MysqlConfig.MaxOpenConns)
	db.SetMaxIdleConns(MysqlConfig.MaxIdleConns)
	if err != nil {

		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = db.Close()
}

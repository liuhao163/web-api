package DbUtils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/beego/bee/logger"
)

var DB *sql.DB

/**

mysql.url=tcp(127.0.0.1:3306)/test?charset=utf8
mysql.username=root
mysql.password=
mysql.maxconns=100
mysql.idelconnes=20
 */
func Init() {
	sqlurl := beego.AppConfig.String("mysql.url")
	username := beego.AppConfig.String("mysql.username")
	password := beego.AppConfig.String("mysql.password")
	maxConns, errorMax := beego.AppConfig.Int("mysql.maxconns")
	idelconnes, errorIdel := beego.AppConfig.Int("mysql.idelconnes")

	if errorMax != nil || errorIdel != nil {
		beeLogger.Log.Error("db param error")
		return
	}

	if len(password) > 0 {
		password = ":" + password
	} else {
		password = "";
	}

	dbUrl := username + password + ":@" + sqlurl
	DB, _ = sql.Open("mysql", dbUrl)
	DB.SetMaxOpenConns(maxConns)
	DB.SetMaxIdleConns(idelconnes)
	//DB.Ping()
}

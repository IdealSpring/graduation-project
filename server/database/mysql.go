package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"graduation-project/server/config"
	"io/ioutil"
	"log"
	"strings"
)

var DB *gorm.DB

// 初始化数据库连接
func init() {
	dbcfg := config.DatabaseConfig
	database := dbcfg.Database
	dbtype := dbcfg.Dbtype
	host := dbcfg.Host
	port := dbcfg.Port
	username := dbcfg.Username
	password := dbcfg.Password

	// 打印完成的数据库URL
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)
	log.Println("database url: ", url)

	// 判断数据库类型并连接数据库
	var db Database
	if strings.ToLower(dbtype) == "mysql" {
		db = new(Mysql)
	} else {
		panic("db type unknow")
	}

	var err error
	DB, err = db.Open(dbtype, url)
	DB.LogMode(true)

	// 错误处理
	if err != nil {
		log.Fatalf("mysql connect error: %v", err)
	} else {
		log.Println("mysql connect success...")
	}

	if DB.Error != nil {
		log.Fatalf("database error: %v", DB.Error)
	}
}

// 数据库接口
type Database interface {
	Open(dbType string, url string) (db *gorm.DB, err error)
}
type Mysql struct {
}
func (*Mysql) Open(dbType string, url string) (db *gorm.DB, err error) {
	db, err = gorm.Open(dbType, url)
	return
}


// 初始化数据库表(删除/创建 表)
func initDB() error {
	sql, err := readSQLForFile("config/db.sql")
	if err != nil {
		fmt.Println("SQL file read failed, err :", err.Error())
		return err
	}

	sqlList := strings.Split(sql, ";")
	for i := 0; i < len(sqlList); i++ {
		if strings.Contains(sqlList[i], "//") {
			log.Println(sqlList[i])
			continue
		}
		sql = strings.Replace(sqlList[i] + ";", "\r\n", "", -1)

		if err = DB.Exec(sql).Error; err != nil {
			if !strings.Contains(err.Error(), "Query was empty") {
				return err
			}
		}
	}

	return nil
}

// 从SQL文件中读取SQL
func readSQLForFile(path string)(string, error) {
	if content, err := ioutil.ReadFile(path); err == nil {
		log.Printf("read sql for %v successful.\n", path)
		log.Println("Use ioutil.ReadFile to read a file:")
		log.Println(string(content))
		return string(content), nil
	} else {
		return "", err
	}
}

// 初始化系统数据库
func InitApplicationDatabase() {
	if !config.ApplicationConfig.IsInit {
		if err := initDB(); err != nil {
			log.Fatalln("database init failed......")
		} else {
			config.SetApplicationDatabaseIsInit()
		}
	}
}

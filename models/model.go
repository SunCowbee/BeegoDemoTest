package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

//表的设计

//定义一个结构体
type User struct {
	Id       int
	Name     string
	PassWord string
	//Pass_Word
}

//在ORM里面__是有特殊含义的

func init() {
	//ORM操作数据库
	//获取连接对象
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")

	//创建表
	orm.RegisterModel(new(User))

	//生成表
	//第一个参数是数据库别名，第二个参数是是否强制更新
	orm.RunSyncdb("default", false, true)
	//操作表

}

//func init() {
//	//连接数据库
//	conn, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
//	if err != nil {
//		beego.Info("链接失败")
//	}
//	defer conn.Close()
//	//建表
//	res, err := conn.Exec("create table user(userName VARCHAR(40),passwd VARCHAR(40))")
//	beego.Info("create table result=", res, err)
//	//插入数据
//	res, err = conn.Exec("insert user(userName,passwd) values(?,?)", "itcast", "heima")
//	beego.Info(res, err)
//	//单行查询
//	row := conn.QueryRow(`select * from user where userName = "wyj"`)
//	var name, pwd string
//	row.Scan(&name, &pwd)
//	beego.Info(name, "------", pwd)
//	//多行查询
//	data, err := conn.Query("SELECT userName from user")
//	var userName string
//	if err == nil {
//		for data.Next() {
//			data.Scan(&userName)
//			beego.Error(userName)
//		}
//	}
//}

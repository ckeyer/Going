package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var CKLog = beego.BeeLogger

func init() {
	CKLog.SetLevel(beego.LevelDebug)
}

type User struct {
	Id    int
	Name  string
	Grade *Grade `orm:"rel(one)"` // OneToOne relation
}

type Grade struct {
	Id    int
	Grade int16
	User  *User `orm:"reverse(one)"` // 设置反向关系(可选)
}

func initORM() {
	// 需要在init中注册定义的model
	// orm.RegisterModel(new(User), new())
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "ckeyer:wangcj@/db_test?charset=utf8", 10, 10)

	orm.RegisterModelWithPrefix("tb_", new(User))
	orm.RegisterModelWithPrefix("tb_", new(Grade))

	name := "default"
	force := true
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		CKLog.Error("create table false")
		return
	}
}

func main() {
	initORM()
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	grade := new(Grade)
	grade.Grade = 30

	user := new(User)
	user.Grade = grade
	user.Name = "slene"

	fmt.Println(o.Insert(grade))
	fmt.Println(o.Insert(user))
}

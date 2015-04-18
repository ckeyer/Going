package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var CKLog = beego.BeeLogger

func init() {
	CKLog.SetLevel(beego.LevelDebug)
}

type User struct {
	Id      int64
	Name    string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

type Blog struct {
	Id      int64
	Content string
	User    *User     `orm:"rel(fk)"` // 设置反向关系(可选)
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func initORM() {
	// 需要在init中注册定义的model
	// orm.RegisterModel(new(User), new())
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "ckeyer:wangcj@/db_test?charset=utf8", 10, 10)

	orm.RegisterModelWithPrefix("tb_", new(User))
	orm.RegisterModelWithPrefix("tb_", new(Blog))

	name := "default"
	force := true
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		CKLog.Error("create table false")
		return
	}
}

func (this *Blog) Update() error {
	o := orm.NewOrm()
	_, e := o.Update(this)
	if e != nil {
		return e
	}

	return nil
}

func main() {
	initORM()
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	user := new(User)
	user.Name = "slene"

	blog := new(Blog)
	blog.User = user
	blog.Content = "hahahhahah"

	// user.Created = time.Now()
	// user.Updated = time.Now()
	// blog.Created = time.Now()
	// blog.Updated = time.Now()

	fmt.Println(o.Insert(user))
	fmt.Println(o.Insert(blog))
	time.Sleep(2 * time.Second)
	blog.Content = "123update"
	blog.Update()
}

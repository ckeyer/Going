package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type TestUser struct {
	Id    int64      `json:"-",xml:"-"`
	Name  string     `json:"name",xml:"Name",orm:"size(16)"`
	Class *TestClass `orm:"one"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}
type TestClass struct {
	Id   int64  `json:"-",xml:"-"`
	Name string `json:"name",xml:"Name",orm:"size(16)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func init(){
	
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "ckeyer:wangcj@/db_test?charset=utf8", 10, 10)

	orm.RegisterModelWithPrefix("tb_", new(TestClass))
	orm.RegisterModelWithPrefix("tb_", new(TestUser))
	
	
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
	class1 := &TestClass{
		Name: "Math",
	}

	user1 := &TestUser{
		Name:  "ckeyer",
		Class: class1,
	}
	bsjson, _ := json.Marshal(user1)
	bsxml, _ := xml.Marshal(user1)

	fmt.Printf("%s\n", bsjson)
	fmt.Printf("%s\n", bsxml)

}


func (this *TestUser) Update() error {
	o := orm.NewOrm()
	_, e := o.Update(this)
	if e != nil {
		return e
	}

	return nil
}

func (this *TestClass) Update() error {
	o := orm.NewOrm()
	_, e := o.Update(this)
	if e != nil {
		return e
	}

	return nil
}

package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
	_ "github.com/go-sql-driver/mysql"
)

var (
	engine *xorm.Engine
)

func init() {
	var err error
	engine,err =  xorm.NewEngine("mysql","root:root@/xorm?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	//添加日志
	engine.SetLogLevel(core.LOG_DEBUG)
	//打印操作SQL
	engine.ShowSQL(true)

	fmt.Println("连接数据库成功。")
}

func main() {
	defer engine.Clone()
	//初始化数据表
	engine.Sync2(new(TSUser))
	//AddUser("小黑","上海市 嘉定区","123@qq.com",26)
	//UpdateUser()
	//DeleteUserById()
	SearchUser()
}

//创建一个结构体，测试Xorm增删改查

type TSUser struct {
	Id int64	`xorm:"pk autoincr"`
	StudentName string `xorm:varchar(100) student_name`
	Address string	`xorm:varchar(500) address`
	Email string `xorm:varchar(100)`
	Age int
}

/**
新增
 */
func AddUser(studentName string,address string,email string,age int) {
	student := TSUser{StudentName:studentName,Address:address,Email:email,Age:age}
	engine.Insert(student)
}

/**
更新
 */
func UpdateUser() {
	student := TSUser{Age:20}
	engine.ID(1).Update(student)
}

/**
删除
 */
func DeleteUserById() {
	student := TSUser{}
	engine.ID(1).Delete(student)
}

/**
查询
 */
func SearchUser() {

	student := &TSUser{}
	engine.ID(2).Get(student)
	fmt.Println(student.StudentName)

}




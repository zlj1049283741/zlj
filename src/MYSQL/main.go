package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
var db *gorm.DB


type User struct {
	Id int
	Name string
	Age int
	Sex string
	Phone string
}
func main() {
	var err error
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	user := User{5,"lisian",19,"女","123234"}
	db.Table("user").Create(user)
	db.Table("user").Where("Name = ? ","lisian").Delete(User{})

	//user1 := User{2,"l二万人an",19,"男","123234"}
	//user1.Insert()
	// user.Insert()
	//db.Table("user").Save(user)
	db.SingularTable(true)

	var t User
	db.Where("Id = ?" , 2).Find(&t)
	fmt.Println(t)
	// now use p

}
func (user *User) Insert()  {
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。
	db.Table("user").Create(user)
}
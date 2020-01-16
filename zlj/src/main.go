package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)
const (
	host     = "192.168.100.172"
	port     = 5432
	user     = "test"
	password = "123456"
	dbname   = "panasonic_demo"
)
func main() {
	show()
	/*
	//启动一个默认的路由
	router := gin.Default()
	//给/配置一个处理函数
	router.LoadHTMLFiles("./login.html")
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",nil)
	})
	//启动webserver
	router.Run(":8080")*/
}
func show()  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	//连接成功
	fmt.Println("Successfully connected!")
}
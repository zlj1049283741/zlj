package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "0.0.0.0"
	port     = 5432
	user     = "test"
	password = "123456"
	dbname   = "panasonic_demo"
)

type User struct {
	id int
	unionId string
	sex string
	telephone int
	email string
	address string
	comAddress string
	departmenr string
}

//传递unionid查询用户信息
func Show(id int) (userinfo User){

	//连接posegresql
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//连接成功
	fmt.Println("Successfully connected!")

	//查询用户信息
	sqlStatement := `SELECT * FROM userinfo WHERE id=$1;`
	//var teacher Teacher
	var user User
	row := db.QueryRow(sqlStatement, id)
	err = row.Scan(&user.id, &user.unionId,&user.sex,&user.telephone,
		&user.email,&user.address,&user.comAddress,&user.departmenr)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		return user
	default:
		panic(err)
	}
}

func main() {
	var user User
	user = Show(1)
	fmt.Println(user)
}

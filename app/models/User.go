package models

import (
	"fmt"
	"slim/app/libs"
)

var Users User

type User struct {
	ID       int    `json:"id" form:"id"`
	UserName string `json:"user_name"`
}

func (user User) getUserList() {
	//users := make([]*User,0)
	DB := libs.ConnectMySQL()
	rows, err := DB.Query("select * from `client`")

	if err != nil {

	}

	for rows.Next() {
		rows.Scan(&user.ID, &user.UserName)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(user.ID, user.UserName)
	defer rows.Close()

}

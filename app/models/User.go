package models

import (
	"fmt"
	"slim/app/libs"
)

type User struct {
	ID     int    `json:"id" form:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Active int    `json:"active"`
}

func (user User) GetUserList() interface{} {
	users := make([]User, 0)
	//users := make([]User, 0)
	DB := libs.ConnectMySQL()
	rows, err := DB.Query("select `id`,`name`,`email`,`active` from `user`")

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Active)
		users = append(users, user)
		//data, err := json.Marshal(user)
		//if err != nil {
		//	panic (err)
		//}
		//fmt.Println(string(data))
	}
	return users
	//fmt.Println(users)
}

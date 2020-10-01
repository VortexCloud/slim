package models

import (
	"fmt"
	"slim/app/libs"
)

type User struct {
	ID       int    `json:"id" form:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-" `
	Active   int    `json:"active"`
}

func (user User) GetUserList() interface{} {
	users := make([]User, 0)
	//users := make([]User, 0)
	db := libs.ConnectMySQL()
	rows, err := db.Query("select `id`,`name`,`email`,`active` from `user`")

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

func (user User) CreateUser(name string, email string, password string) bool {
	db := libs.ConnectMySQL()
	stmt, err := db.Prepare("INSERT `user` SET name=?, email=?, password=?")
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := stmt.Exec(name, email, password)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res)

	return true
}

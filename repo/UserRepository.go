package UserRepository

import (
	"web-api/models/user"
	"web-api/utils"
	"fmt"
)

func Save(user *user.User) (int64, error) {
	res, err := DbUtils.GetDb().Exec("INSERT INTO user(name) VALUES(?)", user.Name)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func GetById(id int64) (user.User, error) {

	var stmt, err = DbUtils.GetDb().Prepare("SELECT * FROM user where  id=?")
	defer stmt.Close()
	if err != nil {
		panic(err)
		return user.User{}, err
	}

	rows, err := stmt.Query(id)
	defer rows.Close()

	if err != nil {
		fmt.Print(err)
		panic(err)
		return user.User{}, err
	}

	if rows.Next() {
		rowsUser := user.User{}
		rows.Scan(&rowsUser.Id, &rowsUser.Name);
		return rowsUser, nil
	}
	return user.User{}, nil
}

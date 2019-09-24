// Author :		Eric<eehsiao@gmail.com>

package main

import (
	"database/sql"
	"fmt"

	mysql "github.com/eehsiao/go-models-mysql"
)

var (
	myDao     *mysql.Dao
	users     []*UserTb
	user      *UserTb
	serialStr string
	keyValues = make(map[string]interface{})
	status    string
	err       error
	redBool   bool
	val       interface{}
	row       *sql.Row
)

func main() {
	myUserDao := &MyUserDao{
		Dao: mysql.NewDao().SetConfig("root", "mYaDmin", "127.0.0.1:3306", "mysql").OpenDB(),
	}
	defer myUserDao.Close()

	// example 1 : use sql builder
	sets := map[string]interface{}{"foo": 1, "bar": "2", "test": true}
	myUserDao.Set(sets).From("user").Where("abc=1")
	fmt.Println("sqlbuilder", myUserDao.BuildUpdateSQL().BuildedSQL())

	// example 2 : directly use the sqlbuilder
	myUserDao.Select("Host", "User", "Select_priv").From("user").Where("User='root'").Limit(1)
	if row, err = myUserDao.GetRow(); err == nil {
		if val, err = myUserDao.ScanRowType(row, (*UserTb)(nil)); err == nil {
			u, _ := val.(*UserTb)
			fmt.Println("UserTb", u)
		}
	}

	// example 3 : use the data logical
	// set a struct for dao as default model (option)
	// (*UserTb)(nil) : nil pointer of the UserTb struct
	// "user" : is real table name in the db
	if err = myUserDao.SetDefaultModel((*UserTb)(nil), "user"); err != nil {
		panic(err.Error())
	}

	if user, err = myUserDao.GetFirstUser(); err == nil {
		fmt.Println("GetFirstUser", user)
	}
}

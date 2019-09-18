# go-models-mysql
`go-models-mysql` its lite and easy model.

## Features
* Field scanning has become easier since the original driver was extended.
Assumption: we have 5 fields to scan
```go
type Tb struct {
	field0 sql.NullString,
	field1 sql.NullString,
	field2 sql.NullString,
	field3 sql.NullString,
	field4 sql.NullString,
}
```

In original driver, it can't dynamic. how many fields, that you must write fields many how. it you have 20 fileds, you must write 20 times.
```go
var tb Tb
err = rows.Scan(&tb.field0, &tb.field1, &tb.field2, &tb.field3, &tb.field4)
```

In go-models-mysql , you just fill struct nil pointer.
```go
if val, err = myDao.ScanRowType(row, (*Tb)(nil)); err == nil {
	u, _ := val.(*Tb)
	fmt.Println("Tb", u)
}
```

* DAO layer let you operate mysql more Intuitively.
	* Original driver (sql.DB) was extended, so you can operate original commands.
		* ex: Query, QueryRow, Exec ....
	* Import the sqlbuilder that help access sql db easily.
	```go
	myDao.Select("Host", "User", "Select_priv").From("user").Where("User='root'").Limit(1)
	```
	* Set the default table in DAO, that you can design your dao layer friendly.
	```go
	// set a struct for dao as default model (option)
	// (*UserTb)(nil) : nil pointer of the UserTb struct
	// "user" : is real table name in the db
	myUserDao.SetDefaultModel((*UserTb)(nil), "user")

	// call model's Get() , get all rows in user table
	// return (rows *sql.Rows, err error)
	rows, err = myDao.Get()
	```

## Requirements
* Go 1.12 or higher.
* [database/sql](https://golang.org/pkg/database/sql/) package
* [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) package

## Go-Module
create `go.mod` file in your package folder, and fill below
```
module github.com/eehsiao/go-models-example

go 1.13

require (
	github.com/eehsiao/go-models-lib latest
	github.com/eehsiao/go-models-mysql latest
	github.com/eehsiao/go-models-redis latest
	github.com/eehsiao/sqlbuilder latest
	github.com/go-sql-driver/mysql v1.4.1
)

```

## Docker
Easy to start the test evn. That you can run the example code.
```bash
$ docker-compose up -d
```

## Usage
```go
import (
    "database/sql"
	"fmt"

	mysql "github.com/eehsiao/go-models-mysql"
	redis "github.com/eehsiao/go-models-redis"
)

// UserTb : sql table struct that to store into mysql
type UserTb struct {
	Host       sql.NullString `TbField:"Host"`
	User       sql.NullString `TbField:"User"`
	SelectPriv sql.NullString `TbField:"Select_priv"`
}

//new mysql dao
myUserDao := &MyUserDao{
    Dao: mysql.NewDao().SetConfig("root", "mYaDmin", "127.0.0.1:3306", "mysql").OpenDB(),
}

// example 1 : directly use the sqlbuilder
myUserDao.Select("Host", "User", "Select_priv").From("user").Where("User='root'").Limit(1)
fmt.Println("sqlbuilder", myUserDao.BuildSelectSQL())
if row, err = myUserDao.GetRow(); err == nil {
    if val, err = myUserDao.ScanRowType(row, (*UserTb)(nil)); err == nil {
        u, _ := val.(*UserTb)
        fmt.Println("UserTb", u)
    }
}
    
// set a struct for dao as default model (option)
// (*UserTb)(nil) : nil pointer of the UserTb struct
// "user" : is real table name in the db
myUserDao.SetDefaultModel((*UserTb)(nil), "user")

// call model's Get() , get all rows in user table
// return (rows *sql.Rows, err error)
rows, err = myDao.Get()

// call model's GetRow() , get first row in user rows
// return (row *sql.Row, err error)
row, err = myDao.GetRow()

```

## Example
### 1 build-in
[example.go](https://github.com/eehsiao/go-models/blob/master/example/example.go)

The example will connect to local mysql and get user data.
Then connect to local redis and set user data, and get back.

### 2 example
`https://github.com/eehsiao/go-models-example/`


## How-to 
How to design model data logical
### MySQL
#### 1.
create a table struct, and add the tag `TbField:"real table filed"`

`TbField` the tag is musted. `read table filed` also be same the table field.
```go
type UserTb struct {
	Host       sql.NullString `TbField:"Host"`
	User       sql.NullString `TbField:"User"`
	SelectPriv sql.NullString `TbField:"Select_priv"`
}
```
#### 2.
use Struce4QuerySlice to gen the sqlbuilder select fields
```go
m := mysql.NewDao().SetConfig("root", "mYaDmin", "127.0.0.1:3306", "mysql").OpenDB()
m.Select(lib.Struce4QuerySlice(m.DaoStructType)...).From(m.TbName).Limit(3)
```
#### 3.
scan the sql result to the struct of object
```go
row, err = m.GetRow()
if val, err = m.ScanRowType(row, (*UserTb)(nil)); err == nil {
    u, _ := val.(*UserTb)
    fmt.Println("UserTb", u)
}
```

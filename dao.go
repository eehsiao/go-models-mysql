// Author :		Eric<eehsiao@gmail.com>

package mysql

import (
	"errors"
	"reflect"

	model "github.com/eehsiao/go-models"
	"github.com/go-sql-driver/mysql"
)

// Dao : the data access object struct
type Dao struct {
	*model.SqlDao
}

// NewDao : create a new empty Dao
func NewDao() *Dao {
	dao := &Dao{}
	dao.SqlDao = model.NewSqlDao()

	return dao

}

// SetDefaultModel : set the struct for this dao default
func (dao *Dao) SetDefaultModel(tb interface{}, deftultTbName string) (err error) {
	structType := reflect.TypeOf(tb).Elem()
	if db == nil || cfg == nil {
		err = errors.New("Do NewConfig() and NewDb() first !!")
	}

	dao.DaoStruct = structType.Name()
	dao.DaoStructType = structType
	dao.SetTbName(deftultTbName)

	return
}

// GetConfig : return mysql.Config
func (dao *Dao) GetConfig() *mysql.Config {
	return getConfig()
}

// SetConfig : set config by user, pw, addr, db
func (dao *Dao) SetConfig(user, pw, addr, db string) *Dao {
	setConfig(user, pw, addr, db)

	return dao
}

// SetOriginConfig : set config by mysql.Config
func (dao *Dao) SetOriginConfig(c *mysql.Config) *Dao {
	setOriginConfig(c)

	return dao
}

// OpenDB : connect to db
func (dao *Dao) OpenDB() *Dao {
	if _, err := openDB(); err != nil {
		dao.PanicOrErrorLog("cannot connect to db")
	}
	dao.DB = db
	dao.SetDbName(getConfig().DBName)

	return dao
}

// OpenDBWithPoolConns : connect to db and set pool conns
func (dao *Dao) OpenDBWithPoolConns(active, idle int) *Dao {
	if _, err := openDBWithPoolConns(active, idle); err != nil {
		dao.PanicOrErrorLog("cannot connect to db")
	}
	return dao

}

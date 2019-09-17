// Author :		Eric<eehsiao@gmail.com>

package mysql

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

var (
	cfg *mysql.Config
	db  *sql.DB
)

// SetConfig : create and set mysql config via go-models
// addr : can with port number. ex: '127.0.0.1:3306', if u want to use default port, just use ip addr ex: '127.0.0.1'
func setConfig(user, pw, addr, dbname string) {
	if cfg != nil {
		panic("already had config !!")
	}

	cfg = mysql.NewConfig()
	cfg.User = user
	cfg.Passwd = pw
	cfg.Addr = addr
	cfg.DBName = dbname

}

// setOriginConfig : set mysql config
func setOriginConfig(c *mysql.Config) {
	if c != nil {
		cfg = c
	}
}

// getConfig : return the config
func getConfig() *mysql.Config {
	return cfg
}

// openDB : open a new mysql connection
func openDB(c ...*mysql.Config) (*sql.DB, error) {
	var err error

	if len(c) > 0 && c[0] != nil {
		cfg = c[0]
	}

	if cfg == nil {
		return nil, errors.New("Do NewConfig() first")
	}

	if db != nil {
		return nil, errors.New("already connect to db")
	}

	if db, err = sql.Open("mysql", cfg.FormatDSN()); err != nil {
		return nil, err
	}

	return db, nil
}

// openDBWithPoolConns : open a new mysql connection and pool conns
func openDBWithPoolConns(active, idle int) (*sql.DB, error) {
	var err error
	if db, err = openDB(); err != nil {
		return nil, err
	}

	// setting connections pool
	db.SetMaxOpenConns(active)
	db.SetMaxIdleConns(idle)

	return db, nil
}

// SetPoolConns : dynamic set  pool conns
func SetPoolConns(active, idle int) (err error) {
	if db == nil {
		err = errors.New("without open db")
		return
	}

	// setting connections pool
	db.SetMaxOpenConns(active)
	db.SetMaxIdleConns(idle)

	return
}

// Tx : the transaction of Dao
type Tx struct {
	*sql.Tx
}

// GetLock : get a session lock via Dao transaction
func (t *Tx) GetLock(key string, secs int) (cnt int, err error) {
	err = t.Tx.QueryRow("SELECT COALESCE(GET_LOCK(?, ?), 0)", key, secs).Scan(&cnt)

	return
}

// ReleaseLock : release a session lock via Dao transaction
func (t *Tx) ReleaseLock(key string) (cnt int, err error) {
	err = t.Tx.QueryRow("SELECT COALESCE(RELEASE_LOCK(?), 0)", key).Scan(&cnt)

	return
}

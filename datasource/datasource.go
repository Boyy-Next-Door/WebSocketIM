package datasource

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	host     = "159.75.5.163:3306"
	username = "root"
	pwd      = "540273945"
	database = "im"
)

//该数据源会在Init.go中得到初始化
var DataSource *sqlx.DB

const InsertMessageRead = "INSERT INTO tb_message(msgId,msgType,contentType,data,fromUid,toUid,createAt,isRead,readAt) VALUES (?,?,?,?,?,?,?,1,?)"
const InsertMessage = "INSERT INTO tb_message(msgId,msgType,contentType,data,fromUid,toUid,createAt,isRead) VALUES (?,?,?,?,?,?,?,0)"
const UpdateMessageStatusToRevoke = "UPDATE tb_message SET isRevoke=1, revokeAt=? WHERE msgId=?"
const UpdateMessageStatusToRead = "UPDATE tb_message SET isRead=1, readAt=? WHERE msgId=?"
const SelectHistory = "select * from " +
	"(SELECT msgId,msgType,data,fromUid,toUid,createAt,isRead,isRevoke from tb_message " +
	"WHERE fromUid=? AND toUid=?  AND createAt < ? ORDER BY createAt LIMIT 0,? ) as alpha " +
	"UNION " +
	"select * from " +
	"(SELECT msgId,msgType,data,fromUid,toUid,createAt,isRead,isRevoke from tb_message " +
	"WHERE fromUid=? AND toUid=? AND createAt < ? ORDER BY createAt LIMIT 0,? ) as beta " +
	"order by createAt limit 0,?"

type Place struct {
}

//在此封装CRUD
func Select(sql string, result interface{}, args ...interface{}) error {
	err := DataSource.Select(result, sql, args...)
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return err
	}
	return nil
}

func Insert(sql string, args ...interface{}) (int, error) {
	ret, err := DataSource.Exec(sql, args...)
	if err != nil {
		return 0, err
	}
	id, err := ret.LastInsertId()
	return int(id), err
}

func InsertTx(tx *sql.Tx, sql string, args ...interface{}) (int, error) {
	ret, err := tx.Exec(sql, args...)
	if err != nil {
		return 0, err
	}
	id, err := ret.LastInsertId()
	return int(id), err
}

func Update(sql string, args ...interface{}) (int, error) {
	ret, err := DataSource.Exec(sql, args...)
	if err != nil {
		return 0, err
	}
	rows, err := ret.RowsAffected()
	return int(rows), err
}

func UpdateTx(tx *sql.Tx, sql string, args ...interface{}) (int, error) {
	ret, err := tx.Exec(sql, args...)
	if err != nil {
		return 0, err
	}
	rows, err := ret.RowsAffected()
	return int(rows), err
}

func InitDB() (err error) {
	dsn := username + ":" + pwd + "@tcp(" + host + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	// 也可以使用MustConnect连接不成功就panic
	DataSource, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return errors.New("initiating database failed.")
	}
	DataSource.SetMaxOpenConns(20)
	DataSource.SetMaxIdleConns(10)
	return nil
}

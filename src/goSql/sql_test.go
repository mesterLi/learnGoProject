package goSql

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"testing"
	"time"
)

const (
	userName = "root"
	password = "Li19980412"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "todo"
)

type Article struct {
	Id int `db:"id" json:"id"`
	Content string `db:"content" json:"content"`
	CreateTime time.Time `db:"create_time" json:"createTime"`
	Author string `db:"author" json:"author"`
	Title string `db:"title" json:"title"`
}

func TestSql(t *testing.T) {
	var article Article
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	t.Log(path)
	db, err := sql.Open("mysql", path)
	if err == nil {
		t.Log("数据库打开成功")
	} else {
		t.Log("数据库打开失败", err)
	}
	defer db.Close()
	rows, err := db.Query("select * from article")
	if err == nil {
		//t.Log(rows.Columns())
		for rows.Next() {
			rows.Scan(&article)
		}
	} else {
		t.Log(err)
	}
	obj, _ := json.Marshal(article)
	t.Log(string(obj))
}

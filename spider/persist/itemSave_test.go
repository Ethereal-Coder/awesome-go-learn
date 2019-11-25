package persist

import (
	"database/sql"
	"fmt"
	"github.com/Ethereal-Coder/awesome-go-learn/spider/engine"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	expected := engine.Item{
		Url:     "http://album.zhenai.com/u/108906739",
		Type:    "HOME",
		Id:      "108906739",
		Payload: "payload",
	}
	const (
		USERNAME = "root"
		PASSWORD = "mysql"
		NETWORK  = "tcp"
		SERVER   = "49.233.168.186"
		PORT     = 32345
		DATABASE = "spider"
	)

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Errorf("Open mysql failed,err:%v\n", err)
		panic(err)
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数

	e := save(DB, expected)
	if e != nil {
		t.Errorf("Item Saver: error saving item %v: %v", expected, e)
	}

	t.Log("finish ...")
}

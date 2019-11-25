package persist

import (
	"database/sql"
	"fmt"
	"github.com/Ethereal-Coder/awesome-go-learn/spider/engine"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = "mysql"
	NETWORK  = "tcp"
	SERVER   = "49.233.168.186"
	PORT     = 32345
	DATABASE = "blog"
)

func ItemSaver() (chan engine.Item, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return nil, err
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			e := save(DB, item)
			if e != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, e)
			}
		}
	}()

	return out, nil
}

func save(DB *sql.DB, item engine.Item) error {
	result, err := DB.Exec("insert INTO items(id,url,type) values(?,?,?)", item.Id, item.Url, item.Type)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v", err)
		return err
	}
	fmt.Println("LastInsertID:", lastInsertID)
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return err
	}
	fmt.Println("RowsAffected:", rowsaffected)
	return nil
}

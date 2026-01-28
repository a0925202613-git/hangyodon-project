package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // 這是連線 PostgreSQL 必備的驅動
)

func main() {
	// 設定連線資訊：請確認 user 改成你的 Mac 帳號 (tosiatung)
	// 如果你當初沒設資料庫密碼，password= 之後可以留空
	connStr := "host=localhost port=5432 user=tosiatung dbname=hangyodon_db sslmode=disable"

	// 1. 嘗試開啟連線
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. 測試是否連線成功
	err = db.Ping()
	if err != nil {
		fmt.Println("❌ 連線失敗，請檢查 pgAdmin 是否開啟或帳號正確")
		log.Fatal(err)
	}

	// 3. 執行查詢：抓取漢頓的資料
	var personality string
	var dream string
	query := "SELECT personality, dream FROM sanrio_characters WHERE name = '人魚漢頓'"

	err = db.QueryRow(query).Scan(&personality, &dream)
	if err != nil {
		log.Fatal("❌ 抓取資料失敗：", err)
	}

	// 4. 印出結果
	fmt.Println("---------------------------------------")
	fmt.Println("🎉 成功從資料庫連線！")
	fmt.Println("角色：人魚漢頓 (Hangyodon)")
	fmt.Println("性格：", personality)
	fmt.Println("夢想：", dream)
	fmt.Println("---------------------------------------")
}

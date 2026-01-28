package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // 這是連線 PostgreSQL 必備的驅動
)

// 新增角色
func addCharacter(db *sql.DB, name string, species string, personality string, dream string) {
	query := `INSERT INTO sanrio_characters(name, species, personality, dream) VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(query, name, species, personality, dream)
	if err != nil {
		log.Fatal("新增失敗：", err)
	}
	fmt.Printf("成功新增角色：%s ！\n", name)
}

// 刪除角色的函數
func deleteCharacter(db *sql.DB, id int) {
	query := `DELETE FROM sanrio_characters WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal("刪除失敗：", err)
	}
	fmt.Printf("🗑️ 已成功刪除 ID 為 %d 的角色\n", id)
}

// 修改角色的函數
func updateCharacter(db *sql.DB, id int, newDream string) {
	query := `UPDATE sanrio_characters SET dream = $1 WHERE id = $2`
	_, err := db.Exec(query, newDream, id)
	if err != nil {
		log.Fatal("修改失敗：", err)
	}
	fmt.Printf("📝 已將 ID %d 的夢想修改為：%s\n", id, newDream)
}

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

	//3.1 新增角色
	addCharacter(db, "Hello Kitty", "貓", "開朗活潑而非常溫柔，偶而有一點點小迷糊", "左耳的紅色蝴蝶結")
	addCharacter(db, "帕恰狗", "狗", "好奇心旺盛", "明明是小狗卻可以兩隻腳走路")

	// 呼叫刪除 (刪除 id 4)
	deleteCharacter(db, 4, 5)

	// 呼叫修改 (修改 id 7)
	updateCharacter(db, 7, "吃無限量的香蕉冰淇淋")

	// 4. 印出結果
	fmt.Println("---------------------------------------")
	fmt.Println("🎉 成功從資料庫連線！")
	fmt.Println("角色：人魚漢頓 (Hangyodon)")
	fmt.Println("性格：", personality)
	fmt.Println("夢想：", dream)
	fmt.Println("---------------------------------------")
	fmt.Println("所有角色新增完畢，快去 pgAdmin 看看吧！")
}

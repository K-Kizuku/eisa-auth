package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:password@localhost:5432/example?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}
	defer db.Close()

	seedFile := "db/sql/seed.sql"
	seedSQL, err := os.ReadFile(seedFile)
	if err != nil {
		log.Fatalf("seed.sqlファイル読み込みエラー: %v", err)
	}

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatalf("トランザクション開始エラー: %v", err)
	}

	_, err = tx.ExecContext(context.Background(), string(seedSQL))
	if err != nil {
		tx.Rollback()
		log.Fatalf("seed.sql実行エラー: %v", err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("トランザクションコミットエラー: %v", err)
	}

	fmt.Println("シードデータの挿入が完了しました。")
}

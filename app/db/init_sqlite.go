// db/init_sqlite.go
package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitTestDB() (*sql.DB, error) {
	// ":memory:" ならメモリ上に作成され、テスト終了後自動消滅
	return sql.Open("sqlite3", ":memory:")
}

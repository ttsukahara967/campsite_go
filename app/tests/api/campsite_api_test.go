package api

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"campsite_go/db"
	"campsite_go/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// テスト用DBの初期化
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	createTable := `
	CREATE TABLE campsites (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		address TEXT,
		description TEXT,
		facilities TEXT,
		price INTEGER,
		image_url TEXT,
		latitude REAL,
		longitude REAL,
		created_at DATETIME
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}
	// 初期データ1件投入
	_, err = db.Exec(`INSERT INTO campsites (name, address, description, facilities, price, image_url, latitude, longitude, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)`,
		"Test Campsite", "Tokyo", "Test Desc", "トイレ, シャワー", 1500, "http://example.com/img.png", 35.0, 139.0)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}
	return db
}

// Ginルーターのセットアップ
func setupRouterWithTestDB(t *testing.T) *gin.Engine {
	testdb := setupTestDB(t)
	dbw := &db.DBWrap{DB: testdb}
	r := gin.Default()
	r.GET("/campsites", handler.ListCampsitesHandler(dbw))
	r.GET("/campsites/:id", handler.GetCampsiteHandler(dbw))
	return r
}

// 一覧APIのテスト
func TestListCampsites(t *testing.T) {
	router := setupRouterWithTestDB(t)

	req, _ := http.NewRequest("GET", "/campsites", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "Test Campsite") {
		t.Errorf("Expected response to contain 'Test Campsite', got: %s", w.Body.String())
	}
}

// 単体取得APIのテスト
func TestGetCampsite(t *testing.T) {
	router := setupRouterWithTestDB(t)

	req, _ := http.NewRequest("GET", "/campsites/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "Test Campsite") {
		t.Errorf("Expected response to contain 'Test Campsite', got: %s", w.Body.String())
	}
}

// 存在しないID取得のテスト
func TestGetCampsite_NotFound(t *testing.T) {
	router := setupRouterWithTestDB(t)

	req, _ := http.NewRequest("GET", "/campsites/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", w.Code)
	}
}

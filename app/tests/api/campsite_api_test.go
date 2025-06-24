package api

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"campsite_go/db"
	"campsite_go/handler"
	"campsite_go/middleware"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/mattn/go-sqlite3"
)

// テスト用JWT生成
func generateTestJWT() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "admin",
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("your_secret_key"))
	return tokenString
}

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
	_, err = db.Exec(`INSERT INTO campsites (name, address, description, facilities, price, image_url, latitude, longitude, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)`,
		"Test Campsite", "Tokyo", "Test Desc", "トイレ, シャワー", 1500, "http://example.com/img.png", 35.0, 139.0)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}
	return db
}

// Ginルーターのセットアップ（JWTミドルウェア有り）
func setupRouterWithTestDB(t *testing.T) *gin.Engine {
	testdb := setupTestDB(t)
	dbw := &db.DBWrap{DB: testdb}
	r := gin.Default()
	r.Use(middleware.JWTAuthMiddleware()) // ★認証必須
	r.GET("/campsites", handler.ListCampsitesHandler(dbw))
	r.GET("/campsites/:id", handler.GetCampsiteHandler(dbw))
	return r
}

// ★ 認証あり 一覧APIのテスト
func TestListCampsites_Authorized(t *testing.T) {
	router := setupRouterWithTestDB(t)
	token := generateTestJWT()
	req, _ := http.NewRequest("GET", "/campsites", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "Test Campsite") {
		t.Errorf("Expected response to contain 'Test Campsite', got: %s", w.Body.String())
	}
}

// ★ 認証なし（失敗） 一覧API
func TestListCampsites_Unauthorized(t *testing.T) {
	router := setupRouterWithTestDB(t)
	req, _ := http.NewRequest("GET", "/campsites", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected 401, got %d", w.Code)
	}
}

// ★ 単体取得API（認証あり）
func TestGetCampsite_Authorized(t *testing.T) {
	router := setupRouterWithTestDB(t)
	token := generateTestJWT()
	req, _ := http.NewRequest("GET", "/campsites/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "Test Campsite") {
		t.Errorf("Expected response to contain 'Test Campsite', got: %s", w.Body.String())
	}
}

// ★ 存在しないID取得（認証あり）
func TestGetCampsite_NotFound_Authorized(t *testing.T) {
	router := setupRouterWithTestDB(t)
	token := generateTestJWT()
	req, _ := http.NewRequest("GET", "/campsites/999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", w.Code)
	}
}

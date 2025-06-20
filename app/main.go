package main

import (
    "log"
    "database/sql"
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
    _ "campsite_go/docs" // go.modのmodule名+docs
)

type DBWrap struct {
    DB *sql.DB
}

// db.goのInitDB関数を使う

func main() {
    db, err := InitDB()
    if err != nil {
        log.Fatalf("DB connection failed: %v", err)
    }
    defer db.Close()
    dbw := &DBWrap{DB: db}

    r := gin.Default()
    // Swagger UIルート
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.GET("/campsites", ListCampsitesHandler(dbw))
    r.GET("/campsites/:id", GetCampsiteHandler(dbw))
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    log.Println("Server running on :8080")
    r.Run(":8080")
}


package main

// @title         Campsite API
// @version       1.0
// @description   This is a sample API for campsites.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

import (
	"campsite_go/db"
	_ "campsite_go/docs" // go.modのmodule名+docs
	"campsite_go/handler"
	"campsite_go/middleware"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db_con, err := db.InitDB()
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer db_con.Close()
	dbw := &db.DBWrap{DB: db_con}

	r := gin.Default()

	// ログイン（トークン発行）は認証不要
	r.POST("/login", handler.LoginHandler)

	// Swagger UIルート
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 以降のAPIにJWT認証をかける
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/campsites", handler.ListCampsitesHandler(dbw))
		auth.GET("/campsites/:id", handler.GetCampsiteHandler(dbw))
		// ほか認証必須API
	}
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Println("Server running on :8080")
	r.Run(":8080")
}

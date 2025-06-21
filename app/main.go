package main

import (
	"campsite_go/db"
	_ "campsite_go/docs" // go.modのmodule名+docs
	"campsite_go/handler"
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
	// Swagger UIルート
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/campsites", handler.ListCampsitesHandler(dbw))
	r.GET("/campsites/:id", handler.GetCampsiteHandler(dbw))
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Println("Server running on :8080")
	r.Run(":8080")
}

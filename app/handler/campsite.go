// handler.go

package handler

import (
	"campsite_go/db"
	"campsite_go/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListCampsitesHandler godoc
// @Summary 全キャンプ場一覧
// @Description キャンプ場を全件取得
// @Tags campsites
// @Produce json
// @Success 200 {array} Campsite
// @Router /campsites [get]
func ListCampsitesHandler(db *db.DBWrap) gin.HandlerFunc {
	return func(c *gin.Context) {
		campsites, err := repository.GetAllCampsites(db.DB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, campsites)
	}
}

// GetCampsiteHandler godoc
// @Summary キャンプ場詳細取得
// @Description ID指定でキャンプ場詳細を取得
// @Tags campsites
// @Produce json
// @Param id path int true "Campsite ID"
// @Success 200 {object} Campsite
// @Failure 404 {object} ErrorResponse
// @Router /campsites/{id} [get]
func GetCampsiteHandler(db *db.DBWrap) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		campsite, err := repository.GetCampsiteByID(db.DB, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusOK, campsite)
	}
}

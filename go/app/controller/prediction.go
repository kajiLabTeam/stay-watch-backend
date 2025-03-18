package controller

import (
	"net/http"
	"strconv"
	"time"

	"Stay_watch/service"

	"github.com/gin-gonic/gin"
)

// 来訪確率の予測
// GET /prediction/visit
func GetVisitPrediction(c *gin.Context) {
	// パラメータの取得
	u := c.QueryArray("user-id")
	w := c.DefaultQuery("weekday", "0")
	t := c.DefaultQuery("time", "23:59")
	i := c.DefaultQuery("is-forward", "true")

	// パラメータの型変換
	if len(u) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameter: user-id is required."})
		return
	}
	var userIDs []int64
	for _, id := range u {
		userId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameter: user-id must be an integer."})
			return
		}
		userIDs = append(userIDs, userId)
	}
	weekday, err := strconv.Atoi(w)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameter: weekday must be an integer."})
		return
	}
	isForward, err := strconv.ParseBool(i)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameter: is-forward must be a boolean."})
		return
	}

	// パラメータのバリデーション
	for _, userId := range userIDs {
		if userId <= 0 {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid query parameter: user-id must be greater than 0."})
			return
		}
	}
	if weekday < 0 || weekday > 6 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid query parameter: weekday must be in 0-6."})
		return
	}
	if _, err := time.Parse("15:04", t); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameter: time must be in the format HH:MM and must be between 00:00 and 23:59."})
		return
	}

	// サービスの呼び出し
	ps := service.PredictionService{}
	predictions, err := ps.GetVisitProbability(userIDs, weekday, t, isForward)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// レスポンスの返却
	c.JSON(http.StatusOK, gin.H{"result": predictions})
}

// 退室確率の予測
// GET /prediction/departure
func GetDeparturePrediction(c *gin.Context) {
	// パラメータの取得
	u := c.QueryArray("user-id")
	w := c.DefaultQuery("weekday", "0")
	t := c.DefaultQuery("time", "23:59")
	i := c.DefaultQuery("is-forward", "true")

	// パラメータの型変換
	if len(u) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameter: user-id is required."})
		return
	}
	var userIDs []int64
	for _, id := range u {
		userId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameter: user-id must be an integer."})
			return
		}
		userIDs = append(userIDs, userId)
	}
	weekday, err := strconv.Atoi(w)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameter: weekday must be an integer."})
		return
	}
	isForward, err := strconv.ParseBool(i)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameter: is-forward must be a boolean."})
		return
	}

	// パラメータのバリデーション
	for _, userId := range userIDs {
		if userId <= 0 {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid query parameter: user-id must be greater than 0."})
			return
		}
	}
	if weekday < 0 || weekday > 6 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid query parameter: weekday must be in 0-6."})
		return
	}
	if _, err := time.Parse("15:04", t); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameter: time must be in the format HH:MM and must be between 00:00 and 23:59."})
		return
	}

	// サービスの呼び出し
	ps := service.PredictionService{}
	predictions, err := ps.GetDepartureProbability(userIDs, weekday, t, isForward)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// レスポンスの返却
	c.JSON(http.StatusOK, gin.H{"result": predictions})
}

package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"Stay_watch/model"
	"Stay_watch/service"
)

// 特定のユーザが特定の時間以降(または'までに')学校に来る(または'帰る')確率を算出
func GetProbability(c *gin.Context) {
	status := c.Param("status") // "reporting" or "leave"
	before := c.Param("before") // "before" or "after"
	user_id := c.Query("user_id")
	str_date := c.Query("date")
	str_time := c.Query("time")

	if user_id == "" || str_date == "" || str_time == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date or time is empty"})
		return
	}

	url := "https://stay-estimate.kajilab.dev/app/probability/" + status + "/" + before + "?user_id=" + user_id + "&date=" + str_date + "&time=" + str_time
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to access the processing server"})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to get probability"})
		return
	}

	body, _ := io.ReadAll(resp.Body)
	var probability model.ProbabilityStayingResponse
	if err := json.Unmarshal(body, &probability); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal probability"})
		return
	}

	c.JSON(http.StatusOK, probability)
}

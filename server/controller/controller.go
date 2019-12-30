package controller

import (
	"fmt"
	"gameLog/database"
	"gameLog/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello worlds",
	})
}

func Login(c *gin.Context) {
	type LoginRes struct {
		UserCode string `json:"user_code"`
	}
	var loginRes LoginRes
	if err := c.ShouldBindJSON(&loginRes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user model.User

	database.Db.FirstOrCreate(&user, model.User{
		UserCode: loginRes.UserCode,
	})

	c.JSON(http.StatusOK, gin.H{
		"code":    "1",
		"data":    user,
		"message": "success",
	})
	return
}

func Report(c *gin.Context) {

	type Data struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
		Correct  bool   `json:"correct"`
		UseTime  int    `json:"use_time"`
	}

	type ReportRes struct {
		UserCode string `json:"user_code"`
		DeviceId string `json:"device_id"`
		GameId   int    `json:"game_id"`
		GameType int    `json:"game_type"`
		BlockId  string `json:"block_id"`
		Data     []Data `json:"data"`
	}

	var reportRes ReportRes

	err := c.BindJSON(&reportRes)
	fmt.Println(reportRes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "-1",
			"message": "缺少参数或者类型错误 " + err.Error(),
		})
		return
	}
	for _, data := range reportRes.Data {
		gameLog := model.Log{
			UserCode: reportRes.UserCode,
			DeviceId: reportRes.DeviceId,
			GameId:   reportRes.GameId,
			GameType: reportRes.GameType,
			BlockId:  reportRes.BlockId,
			Question: data.Question,
			Answer:   data.Answer,
			Correct:  data.Correct,
			UseTime:  data.UseTime,
		}
		database.Db.Create(&gameLog)
	}

	var user model.User

	database.Db.FirstOrCreate(&user, model.User{
		UserCode: reportRes.UserCode,
	})
	if reportRes.GameId == 1 && user.GameOneStatus < reportRes.GameType {
		user.GameOneStatus = reportRes.GameType
	} else if reportRes.GameId == 2 && user.GameTwoStatus < reportRes.GameType {
		user.GameTwoStatus = reportRes.GameType
	}
	database.Db.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"code":    "1",
		"message": "success",
	})
}

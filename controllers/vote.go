package controllers

import (
	"go-ranking/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VoteController struct{}

func (v VoteController) AddVote(c *gin.Context) {
	userIdStr := c.DefaultPostForm("userId", "0")
	playerIdStr := c.DefaultPostForm("playerId", "0")

	userId, _ := strconv.Atoi(userIdStr)
	playerId, _ := strconv.Atoi(playerIdStr)

	if userId == 0 || playerId == 0 {
		RetuenError(c, 4004, "请输入正确信息！")
	}

	user, _ := models.GetUserInfo(userId)
	if user.Id == 0 {
		RetuenError(c, 4001, "投票用户不存在")
		return
	}

	player, _ := models.GetPlayerInfo(playerId)
	if player.Id == 0 {
		RetuenError(c, 4001, "投票用户不存在")
		return
	}

	// 判断是否已经投过
	vote, _ := models.GetVoteInfo(userId, playerId)
	if vote.Id != 0 {
		RetuenError(c, 4001, "请勿重复投票")
	}

	// 进行投票
	res, err := models.AddVote(userId, playerId)
	if err == nil {
		// 给对应的选手加分
		models.UpdatePlayerScore(playerId)
		RetuenSucess(c, 200, "投票成功", res, 1)
		return
	}
	RetuenError(c, 4001, "出现问题,请联系管理员！")
	return
}

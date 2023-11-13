package controllers

import (
	"go-ranking/models"
	"strconv"
	"github.com/gin-gonic/gin"
)

type PlayerController struct{}

func (p PlayerController) GetAllPlayer(c *gin.Context){
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr) 

	players, err := models.GetAllPlayers(aid) 
	if err != nil{
		RetuenError(c, 4004, "没有相关信息")
		return
	}

	RetuenSucess(c, 200, "success", players, int64(len(players)))
}

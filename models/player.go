package models

import (
	"go-ranking/dao"

	"gorm.io/gorm"
)

type Player struct {
	Id          int    `json:"id"`
	Aid         int    `json:"aid"`
	Ref         string `json:"ref"`
	Nickname    string `json:"nickname"`
	Declaration string `json:"declaration"`
	Avatar      string `json:"avatar"`
	Score       int    `json:"score"`
}

func (Player) TableName() string {
	return "player"
}

func GetAllPlayers(aid int) ([]Player, error) {
	var player []Player
	err := dao.DB.Where("aid = ?", aid).Find(&player).Error
	return player, err
}

func GetPlayerInfo(id int)(Player, error) {
	var player Player 
	err := dao.DB.Where("id = ?", id).First(&player).Error
	return player, err
}

func UpdatePlayerScore(id int) {
	var player Player 
	dao.DB.Model(&player).Where("id = ?", id).UpdateColumn("score", gorm.Expr("score + ?", 1))
}

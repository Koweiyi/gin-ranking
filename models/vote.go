package models

import (
	"go-ranking/dao"
	"time"
)

type Vote struct {
	Id       int   `json:"id"`
	UserId   int   `json:"userId"`
	PlayerId int   `josn:"playerId"`
	AddTime  int64 `json:"addTime"`
}

func (Vote) TableName() string{
	return "vote"
}

func GetVoteInfo(userId, playerID int) (Vote, error) {
	var vote Vote
	err := dao.DB.Where("user_id = ? and player_id = ?", userId, playerID).First(&vote).Error
	return vote, err
}

func AddVote(userId, playerId int) (int, error) {
	vote := &Vote{UserId: userId, PlayerId: playerId,AddTime: time.Now().Unix()}
	err := dao.DB.Create(vote).Error
	return vote.Id, err
}
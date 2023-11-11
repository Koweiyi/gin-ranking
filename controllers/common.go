package controllers

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

type SucessJsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

type ErrorJsonStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

func RetuenSucess(ctx *gin.Context, code int, msg string, data interface{}, count int64) {
	json := &SucessJsonStruct{Code: code, Msg: msg, Data: data, Count: count}
	ctx.JSON(200, json)
}

func RetuenError(ctx *gin.Context, code int, msg interface{}) {
	json := &SucessJsonStruct{Code: code, Msg: msg}
	ctx.JSON(200, json)
}

func EncryMd5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}

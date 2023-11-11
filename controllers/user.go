package controllers

import (
	"strconv"

	"go-ranking/models"
	// "example.com/gin-ranking/pkg/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

type UserApi struct {
	Id       int    `json:" id"`
	Username string `json:"username"`
}

// 注册功能 接收用户名 username, 密码 password, 确认密码 confirmPassword 三个参数
func (u UserController) Register(c *gin.Context) {
	// 接收用户名， 密码， 确认密码

	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPWD := c.DefaultPostForm("confirmPassword", "")
	if username == "" || password == "" || confirmPWD == "" {
		RetuenError(c, 4001, "请输入正确的信息")
		return
	}
	if password != confirmPWD {
		RetuenError(c, 4001, "密码与确认密码不相同！")
		return
	}

	user, _ := models.GetUserInfoByUsername(username)
	if user.Id != 0 {
		RetuenError(c, 4001, "用户名已存在")
		return
	}
	_, err := models.AddUser(username, EncryMd5(password))

	if err != nil {
		RetuenError(c, 4001, "保存失败，请联系管理员")
		return
	}

	RetuenSucess(c, 200, "保存成功", "", 1)
}

// 登录功能, 接收 用户名 username, 密码 password 两个参数
func (u UserController) Login(ctx *gin.Context) {
	username := ctx.DefaultPostForm("username", "")
	password := ctx.DefaultPostForm("password", "")

	if username == "" || password == "" {
		RetuenError(ctx, 4001, "请输入正确信息")
		return
	}

	user, err := models.GetUserInfoByUsername(username)
	if user.Id == 0 || err != nil || user.Password != EncryMd5(password) {
		RetuenError(ctx, 4004, "用户名或密码不正确")
		return
	}
	// 保存用户登录信息
	session := sessions.Default(ctx)
	session.Set("login:"+strconv.Itoa(user.Id), user.Id)
	session.Save()

	userapi := UserApi{Id: user.Id, Username: user.Username}
	RetuenSucess(ctx, 200, "登录成功", userapi, 1)
	return
}

func (u UserController) GetUserInfo(ctx *gin.Context) {
	idstr := ctx.Param("id")
	name := ctx.Param("name")
	id, _ := strconv.Atoi(idstr)
	user, _ := models.GetUserTest(id)

	RetuenSucess(ctx, 200, name, user, 1)
}

// func (u UserController) GetList(ctx *gin.Context) {
// 	logger.Write("日志信息", "user")
// 	param := Search{}
// 	if e := ctx.BindJSON(&param); e != nil {
// 		RetuenError(ctx, 4001, gin.H{"err": e})
// 		return
// 	}
// 	RetuenSucess(ctx, 200, param.Name, param.Cid, 1)
// }

func (u UserController) MakeError(ctx *gin.Context) {
	// logger.Write("日志信息", "use")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("捕获异常：", err)
	// 	}
	// }()
	zero := 0
	num1 := 1
	res := num1 / zero
	RetuenError(ctx, 4004, res)
}

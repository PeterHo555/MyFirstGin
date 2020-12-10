package controller

import (
	"ginessential/common"
	"ginessential/dto"
	"ginessential/model"
	"ginessential/response"
	"ginessential/util"
	gin "github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()

	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//fmt.Println(name, " ", telephone, " ", password  )

	// 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"密码不能小于6位"})
		return
	}
	// 如果名称没有传，给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}

	// 创建用户
	// 密码加密
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		//ctx.JSON(http.StatusInternalServerError, gin.H{"code":500, "msg":"加密错误"})
		return
	}
	newUser := model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hasePassword),
	}
	DB.Create(&newUser)

	// 返回结果
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}



func Login(ctx *gin.Context) {
	DB := common.GetDB()

	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")

		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")

		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"密码不能小于6位"})
		return
	}

	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")

		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"用户不存在"})
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码错误")

		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"密码错误"})
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统异常")

		//ctx.JSON(http.StatusInternalServerError, gin.H{"code":500, "msg":"系统异常"})
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	ctx.JSON(200, gin.H{
		"code":200,
		"data":gin.H{"token":token},
		"msg":"登陆成功",
	})
}

func Info(ctx *gin.Context)  {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code":200, "data":gin.H{"user": dto.ToUserDto(user.(model.User)) }}) //user.(dto.UserDto)
}




func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}



